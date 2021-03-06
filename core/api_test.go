/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */
package rocketmq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProducerConfig_String(t *testing.T) {
	pConfig := ProducerConfig{}
	pConfig.GroupID = "testGroup"
	pConfig.NameServer = "localhost:9876"
	pConfig.NameServerDomain = "domain1"
	pConfig.GroupName = "producerGroupName"
	pConfig.InstanceName = "testProducer"
	pConfig.LogC = &LogConfig{
		Path:     "/rocketmq/log",
		FileNum:  16,
		FileSize: 1 << 20,
		Level:    LogLevelDebug}
	pConfig.SendMsgTimeout = 30
	pConfig.CompressLevel = 4
	pConfig.MaxMessageSize = 1024

	expect := "ProducerConfig=[GroupId: testGroup, NameServer: localhost:9876, NameServerDomain: domain1, " +
		"GroupName: producerGroupName, InstanceName: testProducer, " +
		"LogConfig: {Path:/rocketmq/log FileNum:16 FileSize:1048576 Level:Debug}, S" +
		"endMsgTimeout: 30, CompressLevel: 4, MaxMessageSize: 1024, ]"
	assert.Equal(t, expect, pConfig.String())
}

func TestPushConsumerConfig_String(t *testing.T) {
	pcConfig := PushConsumerConfig{}
	pcConfig.GroupID = "testGroup"
	pcConfig.NameServer = "localhost:9876"
	pcConfig.GroupName = "consumerGroupName"
	pcConfig.InstanceName = "testPushConsumer"
	pcConfig.LogC = &LogConfig{
		Path:     "/rocketmq/log",
		FileNum:  16,
		FileSize: 1 << 20,
		Level:    LogLevelDebug}
	pcConfig.ThreadCount = 4
	expect := "PushConsumerConfig=[GroupId: testGroup, NameServer: localhost:9876, " +
		"GroupName: consumerGroupName, InstanceName: testPushConsumer, " +
		"LogConfig: {Path:/rocketmq/log FileNum:16 FileSize:1048576 Level:Debug}, ThreadCount: 4, ]"
	assert.Equal(t, expect, pcConfig.String())

	pcConfig.NameServerDomain = "domain1"
	expect = "PushConsumerConfig=[GroupId: testGroup, NameServer: localhost:9876, NameServerDomain: domain1, " +
		"GroupName: consumerGroupName, InstanceName: testPushConsumer, " +
		"LogConfig: {Path:/rocketmq/log FileNum:16 FileSize:1048576 Level:Debug}, ThreadCount: 4, ]"
	assert.Equal(t, expect, pcConfig.String())

	pcConfig.MessageBatchMaxSize = 32
	pcConfig.Model = Clustering
	expect = "PushConsumerConfig=[GroupId: testGroup, NameServer: localhost:9876, NameServerDomain: domain1, " +
		"GroupName: consumerGroupName, InstanceName: testPushConsumer, " +
		"LogConfig: {Path:/rocketmq/log FileNum:16 FileSize:1048576 Level:Debug}, ThreadCount: 4," +
		" MessageBatchMaxSize: 32, MessageModel: Clustering, ]"
	assert.Equal(t, expect, pcConfig.String())
}
