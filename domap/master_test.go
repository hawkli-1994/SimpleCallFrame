package domap

import (
	"container/list"
	"reflect"
	"sync"
	"testing"
)

func TestHelper(t *testing.T) {
	tests := []struct {
		name string
		want *Master
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Helper(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Helper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaster_GetResults(t *testing.T) {
	type fields struct {
		mu      sync.Mutex
		queue   chan *Task
		stop    chan int
		tasks   *list.List
		results []*Result
		f       Handler
		con     int
		timeout int
	}
	tests := []struct {
		name   string
		fields fields
		want   []*Result
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Master{
				mu:      tt.fields.mu,
				queue:   tt.fields.queue,
				stop:    tt.fields.stop,
				tasks:   tt.fields.tasks,
				results: tt.fields.results,
				f:       tt.fields.f,
				con:     tt.fields.con,
				timeout: tt.fields.timeout,
			}
			if got := m.GetResults(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetResults() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaster_Run(t *testing.T) {
	type fields struct {
		mu      sync.Mutex
		queue   chan *Task
		stop    chan int
		tasks   *list.List
		results []*Result
		f       Handler
		con     int
		timeout int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Master{
				mu:      tt.fields.mu,
				queue:   tt.fields.queue,
				stop:    tt.fields.stop,
				tasks:   tt.fields.tasks,
				results: tt.fields.results,
				f:       tt.fields.f,
				con:     tt.fields.con,
				timeout: tt.fields.timeout,
			}
			if got := m.Run(); got != tt.want {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaster_SetCon(t *testing.T) {
	type fields struct {
		mu      sync.Mutex
		queue   chan *Task
		stop    chan int
		tasks   *list.List
		results []*Result
		f       Handler
		con     int
		timeout int
	}
	type args struct {
		con int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Master
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Master{
				mu:      tt.fields.mu,
				queue:   tt.fields.queue,
				stop:    tt.fields.stop,
				tasks:   tt.fields.tasks,
				results: tt.fields.results,
				f:       tt.fields.f,
				con:     tt.fields.con,
				timeout: tt.fields.timeout,
			}
			if got := m.SetCon(tt.args.con); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaster_SetData(t *testing.T) {
	type fields struct {
		mu      sync.Mutex
		queue   chan *Task
		stop    chan int
		tasks   *list.List
		results []*Result
		f       Handler
		con     int
		timeout int
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Master
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Master{
				mu:      tt.fields.mu,
				queue:   tt.fields.queue,
				stop:    tt.fields.stop,
				tasks:   tt.fields.tasks,
				results: tt.fields.results,
				f:       tt.fields.f,
				con:     tt.fields.con,
				timeout: tt.fields.timeout,
			}
			if got := m.SetData(tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaster_SetFunc(t *testing.T) {
	type fields struct {
		mu      sync.Mutex
		queue   chan *Task
		stop    chan int
		tasks   *list.List
		results []*Result
		f       Handler
		con     int
		timeout int
	}
	type args struct {
		f Handler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Master
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Master{
				mu:      tt.fields.mu,
				queue:   tt.fields.queue,
				stop:    tt.fields.stop,
				tasks:   tt.fields.tasks,
				results: tt.fields.results,
				f:       tt.fields.f,
				con:     tt.fields.con,
				timeout: tt.fields.timeout,
			}
			if got := m.SetFunc(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaster_SetRes(t *testing.T) {
	type fields struct {
		mu      sync.Mutex
		queue   chan *Task
		stop    chan int
		tasks   *list.List
		results []*Result
		f       Handler
		con     int
		timeout int
	}
	type args struct {
		res *Result
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Master{
				mu:      tt.fields.mu,
				queue:   tt.fields.queue,
				stop:    tt.fields.stop,
				tasks:   tt.fields.tasks,
				results: tt.fields.results,
				f:       tt.fields.f,
				con:     tt.fields.con,
				timeout: tt.fields.timeout,
			}
		})
	}
}

func TestMaster_SetTimeout(t *testing.T) {
	type fields struct {
		mu      sync.Mutex
		queue   chan *Task
		stop    chan int
		tasks   *list.List
		results []*Result
		f       Handler
		con     int
		timeout int
	}
	type args struct {
		timeout int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Master
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Master{
				mu:      tt.fields.mu,
				queue:   tt.fields.queue,
				stop:    tt.fields.stop,
				tasks:   tt.fields.tasks,
				results: tt.fields.results,
				f:       tt.fields.f,
				con:     tt.fields.con,
				timeout: tt.fields.timeout,
			}
			if got := m.SetTimeout(tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaster_Stop(t *testing.T) {
	type fields struct {
		mu      sync.Mutex
		queue   chan *Task
		stop    chan int
		tasks   *list.List
		results []*Result
		f       Handler
		con     int
		timeout int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Master{
				mu:      tt.fields.mu,
				queue:   tt.fields.queue,
				stop:    tt.fields.stop,
				tasks:   tt.fields.tasks,
				results: tt.fields.results,
				f:       tt.fields.f,
				con:     tt.fields.con,
				timeout: tt.fields.timeout,
			}
		})
	}
}