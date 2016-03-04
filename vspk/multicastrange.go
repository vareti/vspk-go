/*
  Copyright (c) 2015, Alcatel-Lucent Inc
  All rights reserved.

  Redistribution and use in source and binary forms, with or without
  modification, are permitted provided that the following conditions are met:
      * Redistributions of source code must retain the above copyright
        notice, this list of conditions and the following disclaimer.
      * Redistributions in binary form must reproduce the above copyright
        notice, this list of conditions and the following disclaimer in the
        documentation and/or other materials provided with the distribution.
      * Neither the name of the copyright holder nor the names of its contributors
        may be used to endorse or promote products derived from this software without
        specific prior written permission.

  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
  ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
  WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY
  DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
  (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
  LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
  ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
  (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
  SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// MultiCastRangeIdentity represents the Identity of the object
var MultiCastRangeIdentity = bambou.Identity{
	Name:     "multicastrange",
	Category: "multicastranges",
}

// MultiCastRangesList represents a list of MultiCastRanges
type MultiCastRangesList []*MultiCastRange

// MultiCastRangesAncestor is the interface of an ancestor of a MultiCastRange must implement.
type MultiCastRangesAncestor interface {
	MultiCastRanges(*bambou.FetchingInfo) (MultiCastRangesList, *bambou.Error)
	CreateMultiCastRanges(*MultiCastRange) *bambou.Error
}

// MultiCastRange represents the model of a multicastrange
type MultiCastRange struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	MaxAddress    string `json:"maxAddress,omitempty"`
	MinAddress    string `json:"minAddress,omitempty"`
}

// NewMultiCastRange returns a new *MultiCastRange
func NewMultiCastRange() *MultiCastRange {

	return &MultiCastRange{}
}

// Identity returns the Identity of the object.
func (o *MultiCastRange) Identity() bambou.Identity {

	return MultiCastRangeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *MultiCastRange) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *MultiCastRange) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the MultiCastRange from the server
func (o *MultiCastRange) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the MultiCastRange into the server
func (o *MultiCastRange) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the MultiCastRange from the server
func (o *MultiCastRange) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// EventLogs retrieves the list of child EventLogs of the MultiCastRange
func (o *MultiCastRange) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}

// CreateEventLog creates a new child EventLog under the MultiCastRange
func (o *MultiCastRange) CreateEventLog(child *EventLog) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the MultiCastRange
func (o *MultiCastRange) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the MultiCastRange
func (o *MultiCastRange) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the MultiCastRange
func (o *MultiCastRange) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the MultiCastRange
func (o *MultiCastRange) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}