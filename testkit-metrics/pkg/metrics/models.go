package metrics

import (
	"errors"
	"sync"

	simplego "github.com/talbs1986/simplego/metrics/pkg/metrics"
)

const (
	TestkitTotalLabel = "_total"
)

type TestkitCounter struct {
	mux      sync.Mutex
	ValueMap map[string]float64
}

// NewTestkitCounter creates a new testkit metric counter
func NewTestkitCounter(name string, labels *[]string) *TestkitCounter {
	c := &TestkitCounter{
		ValueMap: map[string]float64{},
		mux:      sync.Mutex{},
	}
	if labels != nil {
		for k := range c.ValueMap {
			c.ValueMap[k] = 0
		}
	}
	c.ValueMap[TestkitTotalLabel] = 0
	return c
}

func (c *TestkitCounter) IncBy(count float64, labels simplego.MetricLabels) error {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k := range c.ValueMap {
		c.ValueMap[k] += count
	}
	c.ValueMap[TestkitTotalLabel] += count
	return nil
}

func (c *TestkitCounter) Inc(labels simplego.MetricLabels) error {
	return c.IncBy(1, labels)
}

type TestkitHistogram struct {
	mux      sync.Mutex
	ValueMap map[string]float64
}

// NewTestkitHistogram creates a new testkit metric histogram - NOT IMPLEMENTED
func NewTestkitHistogram(name string, buckets []float64, labels *[]string) *TestkitHistogram {
	c := &TestkitHistogram{
		ValueMap: map[string]float64{},
		mux:      sync.Mutex{},
	}
	if labels != nil {
		for k := range c.ValueMap {
			c.ValueMap[k] = 0
		}
	}
	c.ValueMap[TestkitTotalLabel] = 0
	return c
}

func (c *TestkitHistogram) Record(count float64, labels simplego.MetricLabels) error {
	return errors.New("not impl")
}

type TestkitGauge struct {
	mux      sync.Mutex
	ValueMap map[string]float64
}

// NewTestkitGauge creates a new testkit metric gauge
func NewTestkitGauge(name string, labels *[]string) *TestkitGauge {
	c := &TestkitGauge{
		ValueMap: map[string]float64{},
		mux:      sync.Mutex{},
	}
	if labels != nil {
		for k := range c.ValueMap {
			c.ValueMap[k] = 0
		}
	}
	c.ValueMap[TestkitTotalLabel] = 0
	return c
}

func (c *TestkitGauge) Set(count float64, labels simplego.MetricLabels) error {
	c.mux.Lock()
	defer c.mux.Unlock()
	if labels != nil {
		for k := range c.ValueMap {
			c.ValueMap[k] = count
		}
	}
	c.ValueMap[TestkitTotalLabel] += count
	return nil
}
func (c *TestkitGauge) IncBy(count float64, labels simplego.MetricLabels) error {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k := range c.ValueMap {
		c.ValueMap[k] += count
	}
	c.ValueMap[TestkitTotalLabel] += count
	return nil
}
func (c *TestkitGauge) Inc(labels simplego.MetricLabels) error {
	return c.IncBy(1, labels)
}
