package stat

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Bandmatch-io/util.bandmatch.io/log"
)

var (
	Address = "localhost"
)

// stat enum
const (
	Signups      = 0
	Logins       = 1
	MessagesSent = 2
	Searches     = 3
	RootViews    = 4
	ProfileViews = 5
	Reports      = 6
	ServerErrors = 7
	UserErrors   = 8
)

var enumStatLookup = map[int]string{
	Signups:      "signups",
	Logins:       "logins",
	MessagesSent: "messagesSent",
	Searches:     "searches",
	RootViews:    "rootViews",
	ProfileViews: "profileViews",
	Reports:      "reports",
	ServerErrors: "serverErrors",
	UserErrors:   "userErrors",
}

var (
	atomicQueue chan int
	timerQueue  chan timeData
	closeQueue  chan struct{}
	done        chan struct{}
)

type timeData struct {
	endpoint string
	ms       int64
}

func Start() {
	atomicQueue = make(chan int, 5)
	timerQueue = make(chan timeData, 5)
	done = make(chan struct{})
	closeQueue = make(chan struct{})
	log.Msgf(log.V, "starting stat service")

	go func() {
		for {
			select {
			case val := <-atomicQueue:
				stat, ok := enumStatLookup[val]
				if !ok {
					log.Msgf(log.VV, "stat [%v] does not exist", val)
				} else {
					log.Msg(log.VV, "sending atomic to server")
					err := makeAtomicRequest(stat)
					if err != nil {
						log.Msgf(log.V, "could not send atomic to server: %v", err)
					}
				}
			case val := <-timerQueue:
				log.Msg(log.VV, "sending time data to server")
				err := makeTimerRequest(val)
				if err != nil {
					log.Msgf(log.V, "could not send time data to server: %v", err)
				}
			case <-closeQueue:
				close(done)
				return
			}
		}
	}()
}

func Close() {
	log.Msgf(log.V, "closing stat service")
	closeQueue <- struct{}{}
	<-done
}

func makeAtomicRequest(stat string) error {
	url := fmt.Sprintf("http://%v/stats/atomic?st=%v", Address, stat)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return fmt.Errorf("could not make post request: %v", err)
	}

	if resp.StatusCode != 202 {
		return fmt.Errorf("did not receive 202 response")
	}

	return nil
}

func makeTimerRequest(stat timeData) error {
	url := fmt.Sprintf("http://%v/stats/timing?ep=%v&tm=%v", Address, stat.endpoint, stat.ms)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return fmt.Errorf("could not make post request: %v", err)
	}

	if resp.StatusCode != 202 {
		return fmt.Errorf("did not receive 202 response")
	}

	return nil
}

// Atomic sends a request to update a statistic
func Atomic(val int) {
	atomicQueue <- val
}

func timer(ep string, val int64) {
	timerQueue <- timeData{ep, val}
}

// RecordTime starts a timer, and returns a function to stop it. Once stopped the timer will
// send time data to the server.
func StartTimer(endpoint string) func() {
	start := time.Now()

	return func() {
		t := time.Now()
		elapsed := t.Sub(start)
		timer(endpoint, elapsed.Milliseconds())
	}
}
