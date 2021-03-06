// +build !linux

/*
 * Minio Client (C) 2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this fs except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import "github.com/rjeczalik/notify"

var (
	// EventTypePut contains the notify events that will cause a put
	EventTypePut = []notify.Event{notify.Create, notify.Write, notify.Rename}
	// EventTypeDelete contains the notify events that will cause a delete
	EventTypeDelete = []notify.Event{notify.Remove}
)

// IsPutEvent checks if the event returned is a put event
func IsPutEvent(event notify.Event) bool {
	switch event {
	case notify.Create:
		return true
	case notify.Rename:
		return true
	case notify.Write:
		return true
	}

	return false
}

// IsDeleteEvent checks if the event returned is a delete event
func IsDeleteEvent(event notify.Event) bool {
	switch event {
	case notify.Remove:
		return true
	}

	return false
}
