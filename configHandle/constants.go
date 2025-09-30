package confighandle

const Version = "v0.0.1"

const AppName = "var-saver"

// Templates for printing to CLI

const varTemplate = `*Variable*
  Name: %s
  Environment: %s
`

const metaTemplate = `---------------------------------------
*Metadata*
Application: %s
Version: %s
Variables Last Updated: %s
---------------------------------------
`

const projectTemplate = `
**Project: '%s'**
Total Variables: %d
Variables:
%s
=======================================
`
