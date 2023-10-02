package acl

	type ACL string

	const (
	 ACLUserGetOne ACL = "auth.User.GetOne"
	     ACLRoleGetOne ACL = "auth.Role.GetOne"
	     ACLroutesAuth ACL = "routes.Auth"
	     ACLServiceMapGetMany ACL = "auth.ServiceMap.GetMany"
	     ACLGeneralGet ACL = "auth.General.Get"
	     ACLGeneralUpdate ACL = "auth.General.Update"
	     ACLAuthLogin ACL = "auth.Auth.Login"
	     ACLServiceMapGetOne ACL = "auth.ServiceMap.GetOne"
	     ACLUserAddMultipleUser ACL = "auth.User.AddMultipleUser"
	     ACLUserGetMany ACL = "auth.User.GetMany"
	     ACLRoleAddRole ACL = "auth.Role.AddRole"
	     ACLAccessAddOne ACL = "auth.Access.AddOne"
	     ACLAuthRegister ACL = "auth.Auth.Register"
	     ACLAuthGetPublicKey ACL = "auth.Auth.GetPublicKey"
	     ACLAccessGetAll ACL = "auth.Access.GetAll"
	     ACLServiceMapAdd ACL = "auth.ServiceMap.Add"
	     ACLRoleGetNames ACL = "auth.Role.GetNames"
	     ACLRoleAddAccess ACL = "auth.Role.AddAccess"
	    )
	