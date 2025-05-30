// Copyright (c) 2016-2024 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testutils

import (
	"github.com/onsi/ginkgo"
	"github.com/sirupsen/logrus"

	"github.com/projectcalico/calico/libcalico-go/lib/logutils"
)

func HookLogrusForGinkgo() {
	// Set up logging formatting.
	logutils.ConfigureFormatter("test")
	logrus.SetOutput(ginkgo.GinkgoWriter)
	logrus.SetLevel(logrus.DebugLevel)
}
