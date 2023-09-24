package acl

type ACL string

const (
   		ACLRoleGetOne ACL = "auth.Role.GetOne"
       ACLAuthLogin ACL = "auth.Auth.Login"
       ACLAuthRegister ACL = "auth.Auth.Register"
       ACLServiceMapAdd ACL = "auth.ServiceMap.Add"
       ACLServiceMapUpdateOne ACL = "auth.ServiceMap.UpdateOne"
       ACLRoleGetNames ACL = "auth.Role.GetNames"
       ACLAccessDeleteOne ACL = "auth.Access.DeleteOne"
       ACLroutesAuth ACL = "routes.Auth"
       ACLServiceMapGetMany ACL = "auth.ServiceMap.GetMany"
       ACLAuthGetPublicKey ACL = "auth.Auth.GetPublicKey"
       ACLAccessAddOne ACL = "auth.Access.AddOne"
       ACLAccessGetAll ACL = "auth.Access.GetAll"
       ACLServiceMapGetOne ACL = "auth.ServiceMap.GetOne"
       ACLServiceMapDeleteOne ACL = "auth.ServiceMap.DeleteOne"
       ACLRoleAddRole ACL = "auth.Role.AddRole"
    )
