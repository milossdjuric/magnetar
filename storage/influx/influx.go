package influx

import (
	"context"
	"errors"
	"fmt"
	mPb "github.com/c12s/scheme/magnetar"
	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)

type Influx struct {
	c client.Client
}

func New(address string) (*Influx, error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: toAddress(address),
	})
	if err != nil {
		return nil, err
	}

	return &Influx{
		c: c,
	}, nil
}

// map[string]map[string]string{
// 	"tags": map[string]string{
// 		"cpu": "cpu-total",
// 		"mem": "mem-total",
// 	},
// 	"cpu": map[string]string{
// 		"idle":   "10.1",
// 		"system": "53.3",
// 		"user":   "46.6",
// 	},
// 	"mem": map[string]string{
// 		"idle":   "10.1",
// 		"system": "53.3",
// 		"user":   "46.6",
// 	},
// 	"points": map[string]string{
// 		"cpu": "cpu_usage",
// 		"mem": "mem_usage",
// 	},
// 	"times": map[string]string{
// 		"cpu": "12355464",
// 		"mem": "56768789",
// 	},
// }
func (i *Influx) Save(ctx context.Context, data *mPb.EventMsg, name, policy string) error {
	id := toID(name)
	// test if metrics db exists, if exist put data, if not crete database
	err := i.Init(ctx, id, policy)
	if err != nil {
		return err
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  id,
		Precision: "s",
		// RetentionPolicy: policy,
	})
	if err != nil {
		return err
	}
	for k, v := range data.Data[TAGS].Data {
		tags := map[string]string{k: v}
		fields := map[string]interface{}{}
		for kk, vv := range data.Data[k].Data {
			fields[kk] = vv
		}

		for _, tv := range data.Data[TIMES].Data {
			pt, err := client.NewPoint(data.Data[POINTS].Data[k], tags, fields, toTime(tv))
			if err != nil {
				return err
			}
			bp.AddPoint(pt)
			fmt.Println("{{WROTE: }}", k)
		}
	}
	// Write the batch
	if err = i.c.Write(bp); err != nil {
		return err
	}
	return nil
}

func (i *Influx) Query(ctx context.Context, query []string, id string) error {
	for _, squery := range query {
		q := client.NewQuery(queryDB(squery), id, "s")
		if response, err := i.c.Query(q); err == nil && response.Error() == nil {
			for _, r := range response.Results {
				for _, _ = range r.Series {
				}
			}
		} else {
			return err
		}
	}
	return errors.New("Provided query do not exists")
}

func (i *Influx) Init(ctx context.Context, id, retention string) error {
	err, exists := i.exists(id)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	q := client.NewQuery(createDB(id), "", "")
	if response, err := i.c.Query(q); err != nil && response.Error() != nil {
		return err
	}

	q = client.NewQuery(createRetentionPolicy(id, retention), "", "")
	if response, err := i.c.Query(q); err != nil && response.Error() != nil {
		return err
	}

	return nil

}

func (i *Influx) exists(id string) (error, bool) {
	q := client.NewQuery(DB_EXISTS, "", "")
	if response, err := i.c.Query(q); err == nil && response.Error() == nil {
		for _, r := range response.Results {
			for _, rz := range r.Series {
				for _, dbs := range rz.Values {
					for _, name := range dbs {
						if name.(string) == id {
							return nil, true
						}
					}
				}
			}
		}
	} else {
		return err, false
	}
	return nil, false
}

func (i *Influx) Close() {
	i.c.Close()
}
