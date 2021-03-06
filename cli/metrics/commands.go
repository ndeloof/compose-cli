/*
   Copyright 2020 Docker Compose CLI authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package metrics

var commandFlags = []string{
	//added to catch scan details
	"--version", "--login",
}

// Generated with generatecommands/main.go
var managementCommands = []string{
	"help",
	"ecs",
	"scan",
	"app",
	"builder",
	"imagetools",
	"buildx",
	"checkpoint",
	"config",
	"container",
	"context",
	"create",
	"image",
	"manifest",
	"network",
	"node",
	"plugin",
	"secret",
	"service",
	"stack",
	"swarm",
	"system",
	"key",
	"signer",
	"trust",
	"volume",
	"login",
	"logout",
	"compose",
}

var commands = []string{
	"bundle",
	"completion",
	"init",
	"inspect",
	"install",
	"deploy",
	"list",
	"ls",
	"cp",
	"merge",
	"pull",
	"push",
	"render",
	"split",
	"status",
	"uninstall",
	"upgrade",
	"validate",
	"version",
	"build",
	"prune",
	"create",
	"bake",
	"f",
	"b",
	"du",
	"rm",
	"stop",
	"use",
	"remove",
	"attach",
	"commit",
	"cp",
	"diff",
	"exec",
	"export",
	"kill",
	"logs",
	"ps",
	"pause",
	"port",
	"rename",
	"restart",
	"run",
	"start",
	"stats",
	"top",
	"unpause",
	"update",
	"wait",
	"aci",
	"ecs",
	"show",
	"history",
	"import",
	"load",
	"images",
	"rmi",
	"save",
	"tag",
	"annotate",
	"connect",
	"disconnect",
	"demote",
	"promote",
	"disable",
	"enable",
	"set",
	"rollback",
	"scale",
	"up",
	"down",
	"services",
	"ca",
	"join",
	"join-token",
	"leave",
	"unlock",
	"unlock-key",
	"df",
	"events",
	"info",
	"generate",
	"add",
	"revoke",
	"sign",
	"login",
	"azure",
	"logout",
	"search",
	"convert",
}
