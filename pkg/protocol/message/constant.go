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
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package message

var MAGIC_CODE_BYTES = [2]byte{0xda, 0xda}

type (
	MessageType      int
	GettyRequestType byte
	GlobalStatus     byte
)

const (
	/**
	 * The constant TYPE_GLOBAL_BEGIN.
	 */
	MessageType_GlobalBegin MessageType = 1
	/**
	 * The constant TYPE_GLOBAL_BEGIN_RESULT.
	 */
	MessageType_GlobalBeginResult MessageType = 2
	/**
	 * The constant TYPE_GLOBAL_COMMIT.
	 */
	MessageType_GlobalCommit MessageType = 7
	/**
	 * The constant TYPE_GLOBAL_COMMIT_RESULT.
	 */
	MessageType_GlobalCommitResult MessageType = 8
	/**
	 * The constant TYPE_GLOBAL_ROLLBACK.
	 */
	MessageType_GlobalRollback MessageType = 9
	/**
	 * The constant TYPE_GLOBAL_ROLLBACK_RESULT.
	 */
	MessageType_GlobalRollbackResult MessageType = 10
	/**
	 * The constant TYPE_GLOBAL_STATUS.
	 */
	MessageType_GlobalStatus MessageType = 15
	/**
	 * The constant TYPE_GLOBAL_STATUS_RESULT.
	 */
	MessageType_GlobalStatusResult MessageType = 16
	/**
	 * The constant TYPE_GLOBAL_REPORT.
	 */
	MessageType_GlobalReport MessageType = 17
	/**
	 * The constant TYPE_GLOBAL_REPORT_RESULT.
	 */
	MessageType_GlobalReportResult MessageType = 18
	/**
	 * The constant TYPE_GLOBAL_LOCK_QUERY.
	 */
	MessageType_GlobalLockQuery MessageType = 21
	/**
	 * The constant TYPE_GLOBAL_LOCK_QUERY_RESULT.
	 */
	MessageType_GlobalLockQueryResult MessageType = 22

	/**
	 * The constant TYPE_BRANCH_COMMIT.
	 */
	MessageType_BranchCommit MessageType = 3
	/**
	 * The constant TYPE_BRANCH_COMMIT_RESULT.
	 */
	MessageType_BranchCommitResult MessageType = 4
	/**
	 * The constant TYPE_BRANCH_ROLLBACK.
	 */
	MessageType_BranchRollback MessageType = 5
	/**
	 * The constant TYPE_BRANCH_ROLLBACK_RESULT.
	 */
	MessageType_BranchRollbackResult MessageType = 6
	/**
	 * The constant TYPE_BRANCH_REGISTER.
	 */
	MessageType_BranchRegister MessageType = 11
	/**
	 * The constant TYPE_BRANCH_REGISTER_RESULT.
	 */
	MessageType_BranchRegisterResult MessageType = 12
	/**
	 * The constant TYPE_BRANCH_STATUS_REPORT.
	 */
	MessageType_BranchStatusReport MessageType = 13
	/**
	 * The constant TYPE_BRANCH_STATUS_REPORT_RESULT.
	 */
	MessageType_BranchStatusReportResult MessageType = 14

	/**
	 * The constant TYPE_SEATA_MERGE.
	 */
	MessageType_SeataMerge MessageType = 59
	/**
	 * The constant TYPE_SEATA_MERGE_RESULT.
	 */
	MessageType_SeataMergeResult MessageType = 60

	/**
	 * The constant TYPE_REG_CLT.
	 */
	MessageType_RegClt MessageType = 101
	/**
	 * The constant TYPE_REG_CLT_RESULT.
	 */
	MessageType_RegCltResult MessageType = 102
	/**
	 * The constant TYPE_REG_RM.
	 */
	MessageType_RegRm MessageType = 103
	/**
	 * The constant TYPE_REG_RM_RESULT.
	 */
	MessageType_RegRmResult MessageType = 104
	/**
	 * The constant TYPE_RM_DELETE_UNDOLOG.
	 */
	MessageType_RmDeleteUndolog MessageType = 111
	/**
	 * the constant TYPE_HEARTBEAT_MSG
	 */
	MessageType_HeartbeatMsg MessageType = 120

	/**
	 * the constant MessageType_BatchResultMsg
	 */
	MessageType_BatchResultMsg MessageType = 121
)

const (
	VERSION = 1

	// MaxFrameLength max frame length
	MaxFrameLength = 8 * 1024 * 1024

	// V1HeadLength v1 head length
	V1HeadLength = 16

	// Request message type
	GettyRequestType_RequestSync GettyRequestType = 0

	// Response message type
	GettyRequestType_Response GettyRequestType = 1

	// Request which no need response
	GettyRequestType_RequestOneway GettyRequestType = 2

	// Heartbeat Request
	GettyRequestType_HeartbeatRequest GettyRequestType = 3

	// Heartbeat Response
	GettyRequestType_HeartbeatResponse GettyRequestType = 4
)

const (

	/**
	 * Un known global status.
	 */
	// Unknown
	GlobalStatusUnKnown GlobalStatus = 0

	/**
	 * The GlobalStatusBegin.
	 */
	// PHASE 1: can accept new branch registering.
	GlobalStatusBegin GlobalStatus = 1

	/**
	 * PHASE 2: Running Status: may be changed any time.
	 */
	// Committing.
	GlobalStatusCommitting GlobalStatus = 2

	/**
	 * The Commit retrying.
	 */
	// Retrying commit after a recoverable failure.
	GlobalStatusCommitRetrying GlobalStatus = 3

	/**
	 * Rollbacking global status.
	 */
	// Rollbacking
	GlobalStatusRollbacking GlobalStatus = 4

	/**
	 * The Rollback retrying.
	 */
	// Retrying rollback after a recoverable failure.
	GlobalStatusRollbackRetrying GlobalStatus = 5

	/**
	 * The Timeout rollbacking.
	 */
	// Rollbacking since timeout
	GlobalStatusTimeoutRollbacking GlobalStatus = 6

	/**
	 * The Timeout rollback retrying.
	 */
	// Retrying rollback  GlobalStatus = since timeout) after a recoverable failure.
	GlobalStatusTimeoutRollbackRetrying GlobalStatus = 7

	/**
	 * All branches can be async committed. The committing is NOT done yet, but it can be seen as committed for TM/RM
	 * client.
	 */
	GlobalStatusAsyncCommitting GlobalStatus = 8

	/**
	 * PHASE 2: Final Status: will NOT change any more.
	 */
	// Finally: global transaction is successfully committed.
	GlobalStatusCommitted GlobalStatus = 9

	/**
	 * The Commit failed.
	 */
	// Finally: failed to commit
	GlobalStatusCommitFailed GlobalStatus = 10

	/**
	 * The Rollbacked.
	 */
	// Finally: global transaction is successfully rollbacked.
	GlobalStatusRollbacked GlobalStatus = 11

	/**
	 * The Rollback failed.
	 */
	// Finally: failed to rollback
	GlobalStatusRollbackFailed GlobalStatus = 12

	/**
	 * The Timeout rollbacked.
	 */
	// Finally: global transaction is successfully rollbacked since timeout.
	GlobalStatusTimeoutRollbacked GlobalStatus = 13

	/**
	 * The Timeout rollback failed.
	 */
	// Finally: failed to rollback since timeout
	GlobalStatusTimeoutRollbackFailed GlobalStatus = 14

	/**
	 * The Finished.
	 */
	// Not managed in session MAP any more
	GlobalStatusFinished GlobalStatus = 15
)
