package syncer

import (
	mPb "github.com/c12s/scheme/magnetar"
)

type Syncer interface {
	Sub(f func(msg *mPb.EventMsg))
}
