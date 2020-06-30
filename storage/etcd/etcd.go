package etcd

import (
	"context"
	"fmt"
	mPb "github.com/c12s/scheme/magnetar"
	"github.com/coreos/etcd/clientv3"
	"github.com/golang/protobuf/proto"
	"regexp"
	"strconv"
	"time"
)

type ETCD struct {
	kv     clientv3.KV
	client *clientv3.Client
}

func New(endpoints []string, timeout time.Duration) (*ETCD, error) {
	cli, err := clientv3.New(clientv3.Config{
		DialTimeout: timeout,
		Endpoints:   endpoints,
	})

	if err != nil {
		return nil, err
	}

	return &ETCD{
		kv:     clientv3.NewKV(cli),
		client: cli,
	}, nil
}

func (e *ETCD) Reserve(ctx context.Context, req *mPb.ReserveMsg) (*mPb.ReserveRsp, error) {
	fmt.Println("{{TOPOLOGY REQUEST}} ", req)
	problems := []string{}
	reserved := int32(0)
	for _, id := range req.Ids {
		parts := split(id) // id is in form regionid.clusterid.nodeid
		key := parts[2]    // nodeid
		nodeKey := nodeKey(key)

		fmt.Println("{{LOOKUP}} ", nodeKey)
		resp, err := e.kv.Get(ctx, nodeKey, clientv3.WithCountOnly())
		if err != nil {
			return nil, err
		}

		if resp.Count == 0 {
			fmt.Println("{{TOPOLOGY}} TRY RESERVE PROBLEM: ", id)
			problems = append(problems, id)
		} else {
			fmt.Println("{{TOPOLOGY}} TRY RESERVE: ", nodeKey)
			// minimum lease TTL is 1200-second reserve node for 20min
			resp, err := e.client.Grant(ctx, 300) // 300 just fo testing
			if err != nil {
				fmt.Println(err.Error())
				return nil, err
			}

			// reserve node for 20min
			reservedKey := reservedKey(key)
			_, err = e.kv.Put(ctx, reservedKey, key, clientv3.WithLease(resp.ID))
			if err != nil {
				return nil, err
			}

			_, err = e.kv.Put(ctx, reservedTTLKey(reservedKey), req.Metricttl, clientv3.WithLease(resp.ID))
			if err != nil {
				return nil, err
			}

			reserved += 1
			fmt.Println("{{TOPOLOGY}} RESERVED: ", key)
		}
	}
	return &mPb.ReserveRsp{
		Taken:    reserved,
		Excluded: problems,
	}, nil
}

func (e *ETCD) List(ctx context.Context, req *mPb.ReserveMsg) (*mPb.ListRsp, error) {
	return nil, nil
}

func (e *ETCD) Free(ctx context.Context, req *mPb.ReserveMsg) (*mPb.ReserveRsp, error) {
	return nil, nil
}

func (e *ETCD) Query(ctx context.Context, req *mPb.DataMsg) (*mPb.ListRsp, error) {
	fmt.Println("{{QUERY}}", req)
	resp, err := e.kv.Get(ctx, query, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	if err != nil {
		return nil, err
	}

	labels := tolabels(req.Data["labels"])
	rez := &mPb.ListRsp{Data: map[string]*mPb.DataMsg{}}
	re := regexp.MustCompile("([0-9]+)([a-z]+)")
	if len(labels) == 0 {
		for _, ev := range resp.Kvs {
			item := &mPb.DataMsg{}
			err = proto.Unmarshal(ev.Value, item)
			if err != nil {
				return nil, err
			}
			rez.Data[extractNodeID(string(ev.Key))] = item
		}
	} else {
		for _, ev := range resp.Kvs {
			item := &mPb.DataMsg{}
			err = proto.Unmarshal(ev.Value, item)
			if err != nil {
				return nil, err
			}

			add := true
			for k, reqVal := range labels {
				if storedVal, ok := item.Data[k]; !ok {
					add = false
					break
				} else {
					if k == "memory" || k == "storage" {
						reS := re.FindStringSubmatch(storedVal)
						rv, err := strconv.ParseInt(reqVal, 10, 64)
						sval, err := strconv.ParseInt(reS[1], 10, 64)
						if err != nil {
							add = false
							break
						}
						sv, err := tobytes(sval, reS[2])
						if err != nil || rv > sv {
							add = false
							break
						}
					} else if k == "cpu" || k == "cores" {
						rv, err := strconv.ParseInt(reqVal, 10, 64)
						sv, err := strconv.ParseInt(storedVal, 10, 64)
						if err != nil || rv > sv {
							add = false
							break
						}
					} else {
						fmt.Println("{{QUERY}}", k, reqVal, storedVal, (reqVal != storedVal))
						if reqVal != storedVal {
							add = false
							break
						}
					}
				}
			}
			if add {
				rez.Data[extractNodeID(string(ev.Key))] = item
			}
		}
	}

	return rez, nil
}

func (e *ETCD) Store(ctx context.Context, key string, msg, labels *mPb.DataMsg) error {
	hostInfo := &mPb.DataMsg{Data: map[string]string{}}
	for k, v := range msg.Data {
		if isHostInfo(key) {
			hostInfo.Data[key] = v
			delete(msg.Data, k)
		}
	}

	// minimum lease TTL is 120-second
	resp, err := e.client.Grant(ctx, 120)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	hostData, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// after 120 seconds, the key will be removed
	_, err = e.client.Put(ctx, key, string(hostData), clientv3.WithLease(resp.ID))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if labels != nil {
		lData, err := proto.Marshal(labels)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		// after 120 seconds, the key will be removed
		_, err = e.client.Put(ctx, labelsKey(key), string(lData), clientv3.WithLease(resp.ID))
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}

	return nil
}

func (e *ETCD) Count(ctx context.Context, reserveKey string) (int64, error) {
	resp, err := e.kv.Get(ctx, reserveKey, clientv3.WithCountOnly())
	if err != nil {
		return 0, err
	}
	return resp.Count, nil
}

func (db *ETCD) Close() { db.client.Close() }
