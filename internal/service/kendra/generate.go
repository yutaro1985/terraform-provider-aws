//go:generate go run ../../generate/tags/main.go -AWSSDKVersion=2 -TagInIDElem=ResourceARN -ListTags -ListTagsInIDElem=ResourceARN -ServiceTagsSlice -UpdateTags -UntagInTagsElem=TagKeys
//go:generate go run ../../generate/servicepackage/main.go
// ONLY generate directives and package declaration! Do not add anything else to this file.

package kendra
