// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package common

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/uber/cherami-server/common/configure"
	"testing"
)

type AuthUtilSuite struct {
	*require.Assertions // override suite.Suite.Assertions with require.Assertions; this means that s.NotNil(nil) will stop the test, not merely log an error
	suite.Suite
}

type serviceConfig struct {
	configure.ServiceConfig
	deploymentName string
}

func (r *serviceConfig) GetDeploymentName() string {
	return r.deploymentName
}

func TestAuthUtilSuite(t *testing.T) {
	suite.Run(t, new(AuthUtilSuite))
}

func (s *AuthUtilSuite) SetupTest() {
	s.Assertions = require.New(s.T()) // Have to define our overridden assertions in the test setup. If we did it earlier, s.T() will return nil
}

func (s *AuthUtilSuite) TestGetResourceURNCreateDestination() {
	mockService := new(MockService)

	config := &serviceConfig{}

	mockService.On("GetConfig").Return(config)

	s.Equal("urn:cherami:dst::", GetResourceURNCreateDestination(mockService, nil))
	s.Equal("urn:cherami:dst::", GetResourceURNCreateDestination(mockService, StringPtr("")))

	config.deploymentName = "zone1"
	s.Equal("urn:cherami:dst:zone1:", GetResourceURNCreateDestination(mockService, nil))
	s.Equal("urn:cherami:dst:zone1:", GetResourceURNCreateDestination(mockService, StringPtr("")))
	s.Equal("urn:cherami:dst:zone1:/", GetResourceURNCreateDestination(mockService, StringPtr("/")))
	s.Equal("urn:cherami:dst:zone1:/", GetResourceURNCreateDestination(mockService, StringPtr("//")))

	config.deploymentName = "Zone2_ABC"
	s.Equal("urn:cherami:dst:zone2_abc:/dst1", GetResourceURNCreateDestination(mockService, StringPtr("/Dst1")))
	s.Equal("urn:cherami:dst:zone2_abc:/root2", GetResourceURNCreateDestination(mockService, StringPtr("/Root2/Dst2")))

	s.Equal("urn:cherami:dst:zone2_abc:dst2", GetResourceURNCreateDestination(mockService, StringPtr("Dst2")))
	s.Equal("urn:cherami:dst:zone2_abc:root2", GetResourceURNCreateDestination(mockService, StringPtr("Root2/Dst2")))
}

func (s *AuthUtilSuite) TestGetResourceURNReadDestination() {
	mockService := new(MockService)

	config := &serviceConfig{}

	mockService.On("GetConfig").Return(config)

	s.Equal("urn:cherami:dst::", GetResourceURNReadDestination(mockService, nil))
	s.Equal("urn:cherami:dst::", GetResourceURNReadDestination(mockService, StringPtr("")))

	config.deploymentName = "zone1"
	s.Equal("urn:cherami:dst:zone1:", GetResourceURNReadDestination(mockService, nil))
	s.Equal("urn:cherami:dst:zone1:", GetResourceURNReadDestination(mockService, StringPtr("")))
	s.Equal("urn:cherami:dst:zone1:/", GetResourceURNReadDestination(mockService, StringPtr("/")))
	s.Equal("urn:cherami:dst:zone1://", GetResourceURNReadDestination(mockService, StringPtr("//")))

	config.deploymentName = "Zone2_ABC"
	s.Equal("urn:cherami:dst:zone2_abc:/dst1", GetResourceURNReadDestination(mockService, StringPtr("/Dst1")))
	s.Equal("urn:cherami:dst:zone2_abc:/root2/dst2", GetResourceURNReadDestination(mockService, StringPtr("/Root2/Dst2")))

	s.Equal("urn:cherami:dst:zone2_abc:dst2", GetResourceURNReadDestination(mockService, StringPtr("Dst2")))
	s.Equal("urn:cherami:dst:zone2_abc:root2/dst2", GetResourceURNReadDestination(mockService, StringPtr("Root2/Dst2")))
}

func (s *AuthUtilSuite) TestGetResourceURNCreateConsumerGroup() {
	mockService := new(MockService)

	config := &serviceConfig{}

	mockService.On("GetConfig").Return(config)

	s.Equal("urn:cherami:cg::", GetResourceURNCreateConsumerGroup(mockService, nil))
	s.Equal("urn:cherami:cg::", GetResourceURNCreateConsumerGroup(mockService, StringPtr("")))

	config.deploymentName = "zone1"
	s.Equal("urn:cherami:cg:zone1:", GetResourceURNCreateConsumerGroup(mockService, nil))
	s.Equal("urn:cherami:cg:zone1:", GetResourceURNCreateConsumerGroup(mockService, StringPtr("")))
	s.Equal("urn:cherami:cg:zone1:/", GetResourceURNCreateConsumerGroup(mockService, StringPtr("/")))
	s.Equal("urn:cherami:cg:zone1:/", GetResourceURNCreateConsumerGroup(mockService, StringPtr("//")))

	config.deploymentName = "Zone2_ABC"
	s.Equal("urn:cherami:cg:zone2_abc:/dst1", GetResourceURNCreateConsumerGroup(mockService, StringPtr("/Dst1")))
	s.Equal("urn:cherami:cg:zone2_abc:/root2", GetResourceURNCreateConsumerGroup(mockService, StringPtr("/Root2/Dst2")))

	s.Equal("urn:cherami:cg:zone2_abc:dst2", GetResourceURNCreateConsumerGroup(mockService, StringPtr("Dst2")))
	s.Equal("urn:cherami:cg:zone2_abc:root2", GetResourceURNCreateConsumerGroup(mockService, StringPtr("Root2/Dst2")))
}
