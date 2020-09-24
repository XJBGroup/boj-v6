package api

type AddPolicyRequest struct {
	Subject string `json:"subject" form:"subject" binding:"required"`
	Object  string `json:"object" form:"object" binding:"required"`
	Action  string `json:"action" form:"action" binding:"required"`
}

type AddPolicyReply struct {
	Code int  `json:"code" form:"code"`
	Data bool `json:"data" form:"data"`
}

type RemovePolicyRequest struct {
	Subject string `json:"subject" form:"subject" binding:"required"`
	Object  string `json:"object" form:"object" binding:"required"`
	Action  string `json:"action" form:"action" binding:"required"`
}

type RemovePolicyReply struct {
	Code int  `json:"code" form:"code"`
	Data bool `json:"data" form:"data"`
}

type HasPolicyRequest struct {
	Subject string `json:"subject" form:"subject" binding:"required"`
	Object  string `json:"object" form:"object" binding:"required"`
	Action  string `json:"action" form:"action" binding:"required"`
}

type HasPolicyReply struct {
	Code int  `json:"code" form:"code"`
	Data bool `json:"data" form:"data"`
}

type AddGroupingPolicyRequest struct {
	Subject string `binding:"required" json:"subject" form:"subject"`
	Group   string `json:"group" form:"group" binding:"required"`
}

type AddGroupingPolicyReply struct {
	Code int  `json:"code" form:"code"`
	Data bool `json:"data" form:"data"`
}

type RemoveGroupingPolicyRequest struct {
	Subject string `json:"subject" form:"subject" binding:"required"`
	Group   string `json:"group" form:"group" binding:"required"`
}

type RemoveGroupingPolicyReply struct {
	Code int  `json:"code" form:"code"`
	Data bool `json:"data" form:"data"`
}

type HasGroupingPolicyRequest struct {
	Subject string `form:"subject" binding:"required" json:"subject"`
	Group   string `json:"group" form:"group" binding:"required"`
}

type HasGroupingPolicyReply struct {
	Code int  `json:"code" form:"code"`
	Data bool `json:"data" form:"data"`
}

func PSerializeAddPolicyRequest(_subject string, _object string, _action string) *AddPolicyRequest {

	return &AddPolicyRequest{
		Subject: _subject,
		Object:  _object,
		Action:  _action,
	}
}
func SerializeAddPolicyRequest(_subject string, _object string, _action string) AddPolicyRequest {

	return AddPolicyRequest{
		Subject: _subject,
		Object:  _object,
		Action:  _action,
	}
}
func _packSerializeAddPolicyRequest(_subject string, _object string, _action string) AddPolicyRequest {

	return AddPolicyRequest{
		Subject: _subject,
		Object:  _object,
		Action:  _action,
	}
}
func PackSerializeAddPolicyRequest(_subject []string, _object []string, _action []string) (pack []AddPolicyRequest) {
	for i := range _subject {
		pack = append(pack, _packSerializeAddPolicyRequest(_subject[i], _object[i], _action[i]))
	}
	return
}
func PSerializeAddPolicyReply(_code int, _data bool) *AddPolicyReply {

	return &AddPolicyReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeAddPolicyReply(_code int, _data bool) AddPolicyReply {

	return AddPolicyReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeAddPolicyReply(_code int, _data bool) AddPolicyReply {

	return AddPolicyReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeAddPolicyReply(_code []int, _data []bool) (pack []AddPolicyReply) {
	for i := range _code {
		pack = append(pack, _packSerializeAddPolicyReply(_code[i], _data[i]))
	}
	return
}
func PSerializeRemovePolicyRequest(_subject string, _object string, _action string) *RemovePolicyRequest {

	return &RemovePolicyRequest{
		Subject: _subject,
		Object:  _object,
		Action:  _action,
	}
}
func SerializeRemovePolicyRequest(_subject string, _object string, _action string) RemovePolicyRequest {

	return RemovePolicyRequest{
		Subject: _subject,
		Object:  _object,
		Action:  _action,
	}
}
func _packSerializeRemovePolicyRequest(_subject string, _object string, _action string) RemovePolicyRequest {

	return RemovePolicyRequest{
		Subject: _subject,
		Object:  _object,
		Action:  _action,
	}
}
func PackSerializeRemovePolicyRequest(_subject []string, _object []string, _action []string) (pack []RemovePolicyRequest) {
	for i := range _subject {
		pack = append(pack, _packSerializeRemovePolicyRequest(_subject[i], _object[i], _action[i]))
	}
	return
}
func PSerializeRemovePolicyReply(_code int, _data bool) *RemovePolicyReply {

	return &RemovePolicyReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeRemovePolicyReply(_code int, _data bool) RemovePolicyReply {

	return RemovePolicyReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeRemovePolicyReply(_code int, _data bool) RemovePolicyReply {

	return RemovePolicyReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeRemovePolicyReply(_code []int, _data []bool) (pack []RemovePolicyReply) {
	for i := range _code {
		pack = append(pack, _packSerializeRemovePolicyReply(_code[i], _data[i]))
	}
	return
}
func PSerializeHasPolicyRequest(_subject string, _object string, _action string) *HasPolicyRequest {

	return &HasPolicyRequest{
		Subject: _subject,
		Object:  _object,
		Action:  _action,
	}
}
func SerializeHasPolicyRequest(_subject string, _object string, _action string) HasPolicyRequest {

	return HasPolicyRequest{
		Subject: _subject,
		Object:  _object,
		Action:  _action,
	}
}
func _packSerializeHasPolicyRequest(_subject string, _object string, _action string) HasPolicyRequest {

	return HasPolicyRequest{
		Subject: _subject,
		Object:  _object,
		Action:  _action,
	}
}
func PackSerializeHasPolicyRequest(_subject []string, _object []string, _action []string) (pack []HasPolicyRequest) {
	for i := range _subject {
		pack = append(pack, _packSerializeHasPolicyRequest(_subject[i], _object[i], _action[i]))
	}
	return
}
func PSerializeHasPolicyReply(_code int, _data bool) *HasPolicyReply {

	return &HasPolicyReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeHasPolicyReply(_code int, _data bool) HasPolicyReply {

	return HasPolicyReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeHasPolicyReply(_code int, _data bool) HasPolicyReply {

	return HasPolicyReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeHasPolicyReply(_code []int, _data []bool) (pack []HasPolicyReply) {
	for i := range _code {
		pack = append(pack, _packSerializeHasPolicyReply(_code[i], _data[i]))
	}
	return
}
func PSerializeAddGroupingPolicyRequest(_subject string, _group string) *AddGroupingPolicyRequest {

	return &AddGroupingPolicyRequest{
		Subject: _subject,
		Group:   _group,
	}
}
func SerializeAddGroupingPolicyRequest(_subject string, _group string) AddGroupingPolicyRequest {

	return AddGroupingPolicyRequest{
		Subject: _subject,
		Group:   _group,
	}
}
func _packSerializeAddGroupingPolicyRequest(_subject string, _group string) AddGroupingPolicyRequest {

	return AddGroupingPolicyRequest{
		Subject: _subject,
		Group:   _group,
	}
}
func PackSerializeAddGroupingPolicyRequest(_subject []string, _group []string) (pack []AddGroupingPolicyRequest) {
	for i := range _subject {
		pack = append(pack, _packSerializeAddGroupingPolicyRequest(_subject[i], _group[i]))
	}
	return
}
func PSerializeAddGroupingPolicyReply(_code int, _data bool) *AddGroupingPolicyReply {

	return &AddGroupingPolicyReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeAddGroupingPolicyReply(_code int, _data bool) AddGroupingPolicyReply {

	return AddGroupingPolicyReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeAddGroupingPolicyReply(_code int, _data bool) AddGroupingPolicyReply {

	return AddGroupingPolicyReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeAddGroupingPolicyReply(_code []int, _data []bool) (pack []AddGroupingPolicyReply) {
	for i := range _code {
		pack = append(pack, _packSerializeAddGroupingPolicyReply(_code[i], _data[i]))
	}
	return
}
func PSerializeRemoveGroupingPolicyRequest(_subject string, _group string) *RemoveGroupingPolicyRequest {

	return &RemoveGroupingPolicyRequest{
		Subject: _subject,
		Group:   _group,
	}
}
func SerializeRemoveGroupingPolicyRequest(_subject string, _group string) RemoveGroupingPolicyRequest {

	return RemoveGroupingPolicyRequest{
		Subject: _subject,
		Group:   _group,
	}
}
func _packSerializeRemoveGroupingPolicyRequest(_subject string, _group string) RemoveGroupingPolicyRequest {

	return RemoveGroupingPolicyRequest{
		Subject: _subject,
		Group:   _group,
	}
}
func PackSerializeRemoveGroupingPolicyRequest(_subject []string, _group []string) (pack []RemoveGroupingPolicyRequest) {
	for i := range _subject {
		pack = append(pack, _packSerializeRemoveGroupingPolicyRequest(_subject[i], _group[i]))
	}
	return
}
func PSerializeRemoveGroupingPolicyReply(_code int, _data bool) *RemoveGroupingPolicyReply {

	return &RemoveGroupingPolicyReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeRemoveGroupingPolicyReply(_code int, _data bool) RemoveGroupingPolicyReply {

	return RemoveGroupingPolicyReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeRemoveGroupingPolicyReply(_code int, _data bool) RemoveGroupingPolicyReply {

	return RemoveGroupingPolicyReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeRemoveGroupingPolicyReply(_code []int, _data []bool) (pack []RemoveGroupingPolicyReply) {
	for i := range _code {
		pack = append(pack, _packSerializeRemoveGroupingPolicyReply(_code[i], _data[i]))
	}
	return
}
func PSerializeHasGroupingPolicyRequest(_subject string, _group string) *HasGroupingPolicyRequest {

	return &HasGroupingPolicyRequest{
		Subject: _subject,
		Group:   _group,
	}
}
func SerializeHasGroupingPolicyRequest(_subject string, _group string) HasGroupingPolicyRequest {

	return HasGroupingPolicyRequest{
		Subject: _subject,
		Group:   _group,
	}
}
func _packSerializeHasGroupingPolicyRequest(_subject string, _group string) HasGroupingPolicyRequest {

	return HasGroupingPolicyRequest{
		Subject: _subject,
		Group:   _group,
	}
}
func PackSerializeHasGroupingPolicyRequest(_subject []string, _group []string) (pack []HasGroupingPolicyRequest) {
	for i := range _subject {
		pack = append(pack, _packSerializeHasGroupingPolicyRequest(_subject[i], _group[i]))
	}
	return
}
func PSerializeHasGroupingPolicyReply(_code int, _data bool) *HasGroupingPolicyReply {

	return &HasGroupingPolicyReply{
		Code: _code,
		Data: _data,
	}
}
func SerializeHasGroupingPolicyReply(_code int, _data bool) HasGroupingPolicyReply {

	return HasGroupingPolicyReply{
		Code: _code,
		Data: _data,
	}
}
func _packSerializeHasGroupingPolicyReply(_code int, _data bool) HasGroupingPolicyReply {

	return HasGroupingPolicyReply{
		Code: _code,
		Data: _data,
	}
}
func PackSerializeHasGroupingPolicyReply(_code []int, _data []bool) (pack []HasGroupingPolicyReply) {
	for i := range _code {
		pack = append(pack, _packSerializeHasGroupingPolicyReply(_code[i], _data[i]))
	}
	return
}
