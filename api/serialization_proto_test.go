// +build proto

/*
Copyright 2015 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api_test

import (
	"encoding/hex"
	"math/rand"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/jchauncey/kubeclient/api"
	apitesting "github.com/jchauncey/kubeclient/api/testing"
	"github.com/jchauncey/kubeclient/api/v1"
	_ "github.com/jchauncey/kubeclient/apis/extensions"
	_ "github.com/jchauncey/kubeclient/apis/extensions/v1beta1"
	"github.com/jchauncey/kubeclient/runtime"
	"github.com/jchauncey/kubeclient/runtime/protobuf"
	"github.com/jchauncey/kubeclient/util"
)

func init() {
	codecsToTest = append(codecsToTest, func(version string, item runtime.Object) (runtime.Codec, error) {
		return protobuf.NewCodec(version, api.Scheme, api.Scheme, api.Scheme), nil
	})
}

func TestProtobufRoundTrip(t *testing.T) {
	obj := &v1.Pod{}
	apitesting.FuzzerFor(t, "v1", rand.NewSource(benchmarkSeed)).Fuzz(obj)
	data, err := obj.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	out := &v1.Pod{}
	if err := out.Unmarshal(data); err != nil {
		t.Fatal(err)
	}
	if !api.Semantic.Equalities.DeepEqual(out, obj) {
		t.Logf("marshal\n%s", hex.Dump(data))
		t.Fatalf("Unmarshal is unequal\n%s", util.ObjectGoPrintSideBySide(out, obj))
	}
}

func BenchmarkEncodeProtobufGeneratedMarshal(b *testing.B) {
	items := benchmarkItems()
	width := len(items)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := items[i%width].Marshal(); err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
}

// BenchmarkDecodeJSON provides a baseline for regular JSON decode performance
func BenchmarkDecodeIntoProtobuf(b *testing.B) {
	items := benchmarkItems()
	width := len(items)
	encoded := make([][]byte, width)
	for i := range items {
		data, err := (&items[i]).Marshal()
		if err != nil {
			b.Fatal(err)
		}
		encoded[i] = data
		validate := &v1.Pod{}
		if err := proto.Unmarshal(data, validate); err != nil {
			b.Fatalf("Failed to unmarshal %d: %v\n%#v", i, err, items[i])
		}
	}

	for i := 0; i < b.N; i++ {
		obj := v1.Pod{}
		if err := proto.Unmarshal(encoded[i%width], &obj); err != nil {
			b.Fatal(err)
		}
	}
}
