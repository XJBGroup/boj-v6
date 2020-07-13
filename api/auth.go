package api

type AddPolicyRequest struct {
}

type AddPolicyReply struct {
	Code int `json:"code" form:"code"`
}

type RemovePolicyRequest struct {
}

type RemovePolicyReply struct {
	Code int `json:"code" form:"code"`
}

type HasPolicyRequest struct {
}

type HasPolicyReply struct {
	Code int `json:"code" form:"code"`
}

type AddGroupingPolicyRequest struct {
}

type AddGroupingPolicyReply struct {
	Code int `json:"code" form:"code"`
}

type RemoveGroupingPolicyRequest struct {
}

type RemoveGroupingPolicyReply struct {
	Code int `json:"code" form:"code"`
}

type HasGroupingPolicyRequest struct {
}

type HasGroupingPolicyReply struct {
	Code int `json:"code" form:"code"`
}

func PSerializeAddPolicyRequest() *AddPolicyRequest {

	return &AddPolicyRequest{}
}
func SerializeAddPolicyRequest() AddPolicyRequest {

	return AddPolicyRequest{}
}
func _packSerializeAddPolicyRequest() AddPolicyRequest {

	return AddPolicyRequest{}
}
func PackSerializeAddPolicyRequest() (pack []AddPolicyRequest) {
	return
}
func PSerializeAddPolicyReply(_code int) *AddPolicyReply {

	return &AddPolicyReply{
		Code: _code,
	}
}
func SerializeAddPolicyReply(_code int) AddPolicyReply {

	return AddPolicyReply{
		Code: _code,
	}
}
func _packSerializeAddPolicyReply(_code int) AddPolicyReply {

	return AddPolicyReply{
		Code: _code,
	}
}
func PackSerializeAddPolicyReply(_code []int) (pack []AddPolicyReply) {
	for i := range _code {
		pack = append(pack, _packSerializeAddPolicyReply(_code[i]))
	}
	return
}
func PSerializeRemovePolicyRequest() *RemovePolicyRequest {

	return &RemovePolicyRequest{}
}
func SerializeRemovePolicyRequest() RemovePolicyRequest {

	return RemovePolicyRequest{}
}
func _packSerializeRemovePolicyRequest() RemovePolicyRequest {

	return RemovePolicyRequest{}
}
func PackSerializeRemovePolicyRequest() (pack []RemovePolicyRequest) {
	return
}
func PSerializeRemovePolicyReply(_code int) *RemovePolicyReply {

	return &RemovePolicyReply{
		Code: _code,
	}
}
func SerializeRemovePolicyReply(_code int) RemovePolicyReply {

	return RemovePolicyReply{
		Code: _code,
	}
}
func _packSerializeRemovePolicyReply(_code int) RemovePolicyReply {

	return RemovePolicyReply{
		Code: _code,
	}
}
func PackSerializeRemovePolicyReply(_code []int) (pack []RemovePolicyReply) {
	for i := range _code {
		pack = append(pack, _packSerializeRemovePolicyReply(_code[i]))
	}
	return
}
func PSerializeHasPolicyRequest() *HasPolicyRequest {

	return &HasPolicyRequest{}
}
func SerializeHasPolicyRequest() HasPolicyRequest {

	return HasPolicyRequest{}
}
func _packSerializeHasPolicyRequest() HasPolicyRequest {

	return HasPolicyRequest{}
}
func PackSerializeHasPolicyRequest() (pack []HasPolicyRequest) {
	return
}
func PSerializeHasPolicyReply(_code int) *HasPolicyReply {

	return &HasPolicyReply{
		Code: _code,
	}
}
func SerializeHasPolicyReply(_code int) HasPolicyReply {

	return HasPolicyReply{
		Code: _code,
	}
}
func _packSerializeHasPolicyReply(_code int) HasPolicyReply {

	return HasPolicyReply{
		Code: _code,
	}
}
func PackSerializeHasPolicyReply(_code []int) (pack []HasPolicyReply) {
	for i := range _code {
		pack = append(pack, _packSerializeHasPolicyReply(_code[i]))
	}
	return
}
func PSerializeAddGroupingPolicyRequest() *AddGroupingPolicyRequest {

	return &AddGroupingPolicyRequest{}
}
func SerializeAddGroupingPolicyRequest() AddGroupingPolicyRequest {

	return AddGroupingPolicyRequest{}
}
func _packSerializeAddGroupingPolicyRequest() AddGroupingPolicyRequest {

	return AddGroupingPolicyRequest{}
}
func PackSerializeAddGroupingPolicyRequest() (pack []AddGroupingPolicyRequest) {
	return
}
func PSerializeAddGroupingPolicyReply(_code int) *AddGroupingPolicyReply {

	return &AddGroupingPolicyReply{
		Code: _code,
	}
}
func SerializeAddGroupingPolicyReply(_code int) AddGroupingPolicyReply {

	return AddGroupingPolicyReply{
		Code: _code,
	}
}
func _packSerializeAddGroupingPolicyReply(_code int) AddGroupingPolicyReply {

	return AddGroupingPolicyReply{
		Code: _code,
	}
}
func PackSerializeAddGroupingPolicyReply(_code []int) (pack []AddGroupingPolicyReply) {
	for i := range _code {
		pack = append(pack, _packSerializeAddGroupingPolicyReply(_code[i]))
	}
	return
}
func PSerializeRemoveGroupingPolicyRequest() *RemoveGroupingPolicyRequest {

	return &RemoveGroupingPolicyRequest{}
}
func SerializeRemoveGroupingPolicyRequest() RemoveGroupingPolicyRequest {

	return RemoveGroupingPolicyRequest{}
}
func _packSerializeRemoveGroupingPolicyRequest() RemoveGroupingPolicyRequest {

	return RemoveGroupingPolicyRequest{}
}
func PackSerializeRemoveGroupingPolicyRequest() (pack []RemoveGroupingPolicyRequest) {
	return
}
func PSerializeRemoveGroupingPolicyReply(_code int) *RemoveGroupingPolicyReply {

	return &RemoveGroupingPolicyReply{
		Code: _code,
	}
}
func SerializeRemoveGroupingPolicyReply(_code int) RemoveGroupingPolicyReply {

	return RemoveGroupingPolicyReply{
		Code: _code,
	}
}
func _packSerializeRemoveGroupingPolicyReply(_code int) RemoveGroupingPolicyReply {

	return RemoveGroupingPolicyReply{
		Code: _code,
	}
}
func PackSerializeRemoveGroupingPolicyReply(_code []int) (pack []RemoveGroupingPolicyReply) {
	for i := range _code {
		pack = append(pack, _packSerializeRemoveGroupingPolicyReply(_code[i]))
	}
	return
}
func PSerializeHasGroupingPolicyRequest() *HasGroupingPolicyRequest {

	return &HasGroupingPolicyRequest{}
}
func SerializeHasGroupingPolicyRequest() HasGroupingPolicyRequest {

	return HasGroupingPolicyRequest{}
}
func _packSerializeHasGroupingPolicyRequest() HasGroupingPolicyRequest {

	return HasGroupingPolicyRequest{}
}
func PackSerializeHasGroupingPolicyRequest() (pack []HasGroupingPolicyRequest) {
	return
}
func PSerializeHasGroupingPolicyReply(_code int) *HasGroupingPolicyReply {

	return &HasGroupingPolicyReply{
		Code: _code,
	}
}
func SerializeHasGroupingPolicyReply(_code int) HasGroupingPolicyReply {

	return HasGroupingPolicyReply{
		Code: _code,
	}
}
func _packSerializeHasGroupingPolicyReply(_code int) HasGroupingPolicyReply {

	return HasGroupingPolicyReply{
		Code: _code,
	}
}
func PackSerializeHasGroupingPolicyReply(_code []int) (pack []HasGroupingPolicyReply) {
	for i := range _code {
		pack = append(pack, _packSerializeHasGroupingPolicyReply(_code[i]))
	}
	return
}
