//Copyright 2019 Expedia, Inc.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package warmup

import (
	"log"
	"mittens/pkg/grpc"
	"mittens/pkg/http"
	"sync"
	"time"
)

// Warmup holds any information needed for the workers to send requests.
type Warmup struct {
	Target             Target
	MaxDurationSeconds int
	Concurrency        int
}

// HTTPWarmupWorker sends HTTP requests to the target using goroutines.
func (w Warmup) HTTPWarmupWorker(wg *sync.WaitGroup, requests <-chan http.Request, headers map[string]string, requestDelayMilliseconds int, requestsSentCounter *int) {
	for request := range requests {
		time.Sleep(time.Duration(requestDelayMilliseconds) * time.Millisecond)

		resp := w.Target.httpClient.SendRequest(request.Method, request.Path, headers, request.Body)

		if resp.Err != nil {
			log.Printf("🔴 Error in request for %s: %v", request.Path, resp.Err)
		} else {
			*requestsSentCounter++

			if resp.StatusCode/100 == 2 {
				log.Printf("%s response for %s %d ms: %v", resp.Type, request.Path, resp.Duration/time.Millisecond, resp.StatusCode)
			} else {
				log.Printf("🔴 %s response for %s %d ms: %v", resp.Type, request.Path, resp.Duration/time.Millisecond, resp.StatusCode)
			}
		}
	}
	wg.Done()
}

// GrpcWarmupWorker sends gRPC requests to the target using goroutines.
func (w Warmup) GrpcWarmupWorker(wg *sync.WaitGroup, requests <-chan grpc.Request, headers []string, requestDelayMilliseconds int, requestsSentCounter *int) {
	for request := range requests {
		time.Sleep(time.Duration(requestDelayMilliseconds) * time.Millisecond)

		resp := w.Target.grpcClient.SendRequest(request.ServiceMethod, request.Message, headers)

		if resp.Err != nil {
			log.Printf("🔴 Error in request for %s: %v", request.ServiceMethod, resp.Err)
		} else {
			*requestsSentCounter++

			log.Printf("%s response for %s %d ms", resp.Type, request.ServiceMethod, resp.Duration/time.Millisecond)
		}

	}
	wg.Done()
}
