package acl

type ACL string

const (
	ACLAccessAddOne        ACL = "auth.Access.AddOne"
	ACLServiceMapAdd       ACL = "auth.ServiceMap.Add"
	ACLServiceMapDeleteOne ACL = "auth.ServiceMap.DeleteOne"
	ACLGeneralGet          ACL = "auth.General.Get"
	ACLUserAddMultipleUser ACL = "auth.User.AddMultipleUser"
	ACLUserGetOne          ACL = "auth.User.GetOne"
	ACLUserGetMany         ACL = "auth.User.GetMany"
	ACLRoleGetOne          ACL = "auth.Role.GetOne"
	ACLAuthRegister        ACL = "auth.Auth.Register"
	ACLServiceMapUpdateOne ACL = "auth.ServiceMap.UpdateOne"
	ACLGeneralUpdate       ACL = "auth.General.Update"
	ACLAuthGetPublicKey    ACL = "auth.Auth.GetPublicKey"
	ACLroutesAuth          ACL = "routes.Auth"
	ACLServiceMapGetOne    ACL = "auth.ServiceMap.GetOne"
	ACLRoleAddRole         ACL = "auth.Role.AddRole"
	ACLRoleAddAccess       ACL = "auth.Role.AddAccess"
	ACLAccessGetMany       ACL = "auth.Access.GetMany"
	ACLRoleGetNames        ACL = "auth.Role.GetNames"
	ACLAuthLogin           ACL = "auth.Auth.Login"
	ACLServiceMapGetMany   ACL = "auth.ServiceMap.GetMany"

	ACLTestEndpoint ACL = "auth.test.endpoint"
)
