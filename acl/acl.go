package acl

	type ACL string

	const (
	 ACLroutesAuth ACL = "routes.Auth"
	     ACLUserAddUser ACL = "auth.User.AddUser"
	     ACLRoleGetNames ACL = "auth.Role.GetNames"
	     ACLRoleDeleteOne ACL = "auth.Role.DeleteOne"
	     ACLRoleAddRole ACL = "auth.Role.AddRole"
	     ACLServiceMapGetOne ACL = "auth.ServiceMap.GetOne"
	     ACLGeneralGet ACL = "auth.General.Get"
	     ACLUserAddMultipleUser ACL = "auth.User.AddMultipleUser"
	     ACLUserGetMany ACL = "auth.User.GetMany"
	     ACLUserGetOne ACL = "auth.User.GetOne"
	     ACLServiceMapGetMany ACL = "auth.ServiceMap.GetMany"
	     ACLServiceMapUpdateOne ACL = "auth.ServiceMap.UpdateOne"
	     ACLGeneralGetMany ACL = "auth.General.GetMany"
	     ACLAuthLogin ACL = "auth.Auth.Login"
	     ACLRoleUpdateAccesses ACL = "auth.Role.UpdateAccesses"
	     ACLRoleRemoveOneAccess ACL = "auth.Role.RemoveOneAccess"
	     ACLAccessUpdateOneName ACL = "auth.Access.UpdateOneName"
	     ACLAccessAddOne ACL = "auth.Access.AddOne"
	     ACLServiceMapAdd ACL = "auth.ServiceMap.Add"
	     ACLGeneralUpdateDev ACL = "auth.General.UpdateDev"
	     ACLRoleGetOne ACL = "auth.Role.GetOne"
	     ACLAuthRegister ACL = "auth.Auth.Register"
	     ACLAuthOtp ACL = "auth.Auth.Otp"
	     ACLAccessLinkedRoles ACL = "auth.Access.LinkedRoles"
	     ACLServiceMapDeleteOne ACL = "auth.ServiceMap.DeleteOne"
	     ACLGeneralUpdateAdmin ACL = "auth.General.UpdateAdmin"
	     ACLAuthGetPublicKey ACL = "auth.Auth.GetPublicKey"
	     ACLUserUpdateUserRole ACL = "auth.User.UpdateUserRole"
	     ACLAccessGetMany ACL = "auth.Access.GetMany"
	    )
	