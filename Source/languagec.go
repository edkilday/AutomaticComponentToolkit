/*++

Copyright (C) 2018 Autodesk Inc. (Original Author)

All rights reserved.

Redistribution and use in source and binary forms, with or without modification,
are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice,
this list of conditions and the following disclaimer in the documentation
and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

--*/

//////////////////////////////////////////////////////////////////////////////////////////////////////
// languagec.go
// functions to generate the C-layer of a library's API (can be used in bindings or implementation)
//////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"path"
	"strings"
	"log"
)

// BuildBindingC builds C-bindings of a library's API in form of automatically generated C functions
func BuildBindingC(component ComponentDefinition, outputFolderBindingC string) error {
	CTypesHeaderName := path.Join(outputFolderBindingC, component.BaseName + "_types.h")
	log.Printf("Creating \"%s\"", CTypesHeaderName)
	err := CreateCTypesHeader(component, CTypesHeaderName)
	if (err != nil) {
		return err;
	}

	CHeaderName := path.Join(outputFolderBindingC, component.BaseName + ".h")
	log.Printf("Creating \"%s\"", CTypesHeaderName)
	err = CreateCAbiHeader(component, CHeaderName)
	if (err != nil) {
		return err
	}

	return nil
}

// CreateCTypesHeader creates a C header file for the types in component's API
func CreateCTypesHeader (component ComponentDefinition, CTypesHeaderName string) (error) {
	hTypesFile, err := CreateLanguageFile(CTypesHeaderName, "  ");
	if (err != nil) {
		return err;
	}
	hTypesFile.WriteCLicenseHeader (component,
		fmt.Sprintf ("This is an autogenerated plain C Header file with basic types in\norder to allow an easy use of %s", component.LibraryName),
		true);

	err = buildCCPPTypesHeader(component, hTypesFile, component.NameSpace, false);
	return err;
}



func buildSharedCCPPTypesHeader(component ComponentDefinition, w LanguageWriter, NameSpace string) (error) {
	w.Writeln("/*************************************************************************************************************************");
	w.Writeln(" Scalar types definition");
	w.Writeln("**************************************************************************************************************************/")
	w.Writeln("")
	w.Writeln("#ifdef %s_USELEGACYINTEGERTYPES", strings.ToUpper(NameSpace));
	w.Writeln("")
	w.Writeln("typedef unsigned char %s_uint8;", NameSpace);
	w.Writeln("typedef unsigned short %s_uint16 ;", NameSpace);
	w.Writeln("typedef unsigned int %s_uint32;", NameSpace);
	w.Writeln("typedef unsigned long long %s_uint64;", NameSpace);
	w.Writeln("typedef char %s_int8;", NameSpace);
	w.Writeln("typedef short %s_int16;", NameSpace);
	w.Writeln("typedef int %s_int32;", NameSpace);
	w.Writeln("typedef long long %s_int64;", NameSpace);
	w.Writeln("")
	w.Writeln("#else // %s_USELEGACYINTEGERTYPES", strings.ToUpper(NameSpace));
	w.Writeln("")
	w.Writeln("#include <stdint.h>")
	w.Writeln("")
	w.Writeln("typedef uint8_t %s_uint8;", NameSpace);
	w.Writeln("typedef uint16_t %s_uint16;", NameSpace);
	w.Writeln("typedef uint32_t %s_uint32;", NameSpace);
	w.Writeln("typedef uint64_t %s_uint64;", NameSpace);
	w.Writeln("typedef int8_t %s_int8;", NameSpace);
	w.Writeln("typedef int16_t %s_int16;", NameSpace);
	w.Writeln("typedef int32_t %s_int32;", NameSpace);
	w.Writeln("typedef int64_t %s_int64 ;", NameSpace);
	w.Writeln("")
	w.Writeln("#endif // %s_USELEGACYINTEGERTYPES", strings.ToUpper(NameSpace));
	w.Writeln("")
	w.Writeln("typedef float %s_single;", NameSpace);
	w.Writeln("typedef double %s_double;", NameSpace);
	w.Writeln("")


	w.Writeln("/*************************************************************************************************************************");
	w.Writeln(" General type definitions");
	w.Writeln("**************************************************************************************************************************/");

	w.Writeln("");
	w.Writeln("typedef %s_int32 %sResult;", NameSpace, NameSpace);
	w.Writeln("typedef void * %sHandle;", NameSpace);
	w.Writeln("typedef void * %s_pvoid;", NameSpace);
	
	w.Writeln("");
	w.Writeln("/*************************************************************************************************************************");
	w.Writeln(" Version for %s", NameSpace);
	w.Writeln("**************************************************************************************************************************/");
	w.Writeln("");
	w.Writeln("#define %s_VERSION_MAJOR %d", strings.ToUpper(NameSpace), majorVersion(component.Version));
	w.Writeln("#define %s_VERSION_MINOR %d", strings.ToUpper(NameSpace), minorVersion(component.Version));
	w.Writeln("#define %s_VERSION_MICRO %d", strings.ToUpper(NameSpace), microVersion(component.Version));
	w.Writeln("#define %s_VERSION_PRERELEASEINFO \"%s\"", strings.ToUpper(NameSpace), preReleaseInfo(component.Version));
	w.Writeln("#define %s_VERSION_BUILDINFO \"%s\"", strings.ToUpper(NameSpace), buildInfo(component.Version));

	w.Writeln("");

	w.Writeln("/*************************************************************************************************************************");
	w.Writeln(" Error constants for %s", NameSpace);
	w.Writeln("**************************************************************************************************************************/");
	w.Writeln("");
	w.Writeln("#define %s_SUCCESS 0", strings.ToUpper (NameSpace));
	for i := 0; i < len(component.Errors.Errors); i++ {
		errorcode := component.Errors.Errors[i];
		if (errorcode.Description != "") {
			w.Writeln("#define %s_ERROR_%s %d /** %s */", strings.ToUpper (NameSpace), errorcode.Name, errorcode.Code, errorcode.Description);
		} else {
			w.Writeln("#define %s_ERROR_%s %d", strings.ToUpper (NameSpace), errorcode.Name, errorcode.Code);
		}
	}

	w.Writeln("");

	w.Writeln("/*************************************************************************************************************************");
	w.Writeln(" Error strings for %s", NameSpace);
	w.Writeln("**************************************************************************************************************************/");
	w.Writeln("");
	w.Writeln("inline const char * %s_GETERRORSTRING (%sResult nErrorCode) {", strings.ToUpper (NameSpace), NameSpace);
	w.Writeln("  switch (nErrorCode) {");
	w.Writeln("    case %s_SUCCESS: return \"no error\";", strings.ToUpper (NameSpace));
	for i := 0; i < len(component.Errors.Errors); i++ {
		errorcode := component.Errors.Errors[i];
		w.Writeln("    case %s_ERROR_%s: return \"%s\";", strings.ToUpper (NameSpace), errorcode.Name, errorcode.Description);
	}
	w.Writeln("    default: return \"unknown error\";");
	w.Writeln("  }");
	w.Writeln("}");
	w.Writeln("");
	
	w.Writeln("/*************************************************************************************************************************");
	w.Writeln(" Declaration of handle classes ");
	w.Writeln("**************************************************************************************************************************/");
	w.Writeln("");
	
	for i := 0; i < len(component.Classes); i++ {
		class := component.Classes[i];
		w.Writeln("typedef %sHandle %s_%s;", NameSpace, NameSpace, class.ClassName);
	}
	w.Writeln("");
	
	return nil
}

func getCMemberLine(member ComponentDefinitionMember, NameSpace string, arraysuffix string, structName string) (string, error) {
	switch (member.Type) {
		case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "single", "double", "bool", "pointer":
			typeName, err := getCParameterTypeName(member.Type, NameSpace, "")
			if (err != nil) {
				return "", err
			}
			return fmt.Sprintf("%s m_%s%s;", typeName, member.Name, arraysuffix), nil
		case "enum":
			return fmt.Sprintf("structEnum%s%s m_%s%s;", NameSpace, member.Class, member.Name, arraysuffix), nil
		default:
			return "", fmt.Errorf ("it is not possible for struct %s to contain a %s member", structName, member.Type);
	}
}

func buildCCPPTypesHeader(component ComponentDefinition, w LanguageWriter, NameSpace string, useCPPTypes bool) (error) {
	sIncludeGuard :=  "__"+strings.ToUpper(NameSpace)+"_TYPES_HEADER"
	if useCPPTypes {
		sIncludeGuard += "_CPP"
	}
	
	w.Writeln("#ifndef %s", sIncludeGuard);
	w.Writeln("#define %s", sIncludeGuard);
	w.Writeln("");

	if (!useCPPTypes) {
		w.Writeln("#include <stdbool.h>");
	}
	w.Writeln("");

	err := buildSharedCCPPTypesHeader(component, w, NameSpace)
	if (err != nil) {
		return err
	}
	
	if useCPPTypes {
		w.Writeln("namespace %s {", NameSpace);
		w.Writeln("");
		w.AddIndentationLevel(1);
	}

	err = buildCCPPEnums(component, w, NameSpace, useCPPTypes)
	if (err!=nil) {
		return err
	}
	err = buildCCPPStructs(component, w, NameSpace, useCPPTypes)
	if (err!=nil) {
		return err
	}
	err = buildCCPPFunctionPointers(component, w, NameSpace, useCPPTypes)
	if (err != nil) {
		return err
	}
	
	if (useCPPTypes) {
		w.AddIndentationLevel(-1);
		w.Writeln("} // namespace %s;", NameSpace);
		w.Writeln("")
		w.Writeln("// define legacy C-names for enums, structs and function types")
		for i := 0; i < len(component.Enums); i++ {
			enum := component.Enums[i];
			w.Writeln("typedef %s::e%s e%s%s;", NameSpace, enum.Name, NameSpace, enum.Name);
		}
		for i := 0; i < len(component.Structs); i++ {
			structinfo := component.Structs[i];
			w.Writeln("typedef %s::s%s s%s%s;", NameSpace, structinfo.Name, NameSpace, structinfo.Name);
		}
		for i := 0; i < len(component.Functions); i++ {
			functiontype := component.Functions[i]
			w.Writeln("typedef %s::%s %s%s;", NameSpace, functiontype.FunctionName, NameSpace, functiontype.FunctionName);
		}
	}

	w.Writeln("");
	w.Writeln("#endif // %s", sIncludeGuard);

	return nil;
}

// GetCMemberDefaultValue returns the defailt value of a member in C-based-languages
func GetCMemberDefaultValue(memberType string, memberClass string, NameSpace string) (string, error) {
	switch (memberType) {
		case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64":
			return "0", nil;
		case "bool":
			return "false", nil;
		case "single":
			return "0.0f", nil;
		case "double":
			return "0.0", nil;
		case "pointer":
			return "nullptr", nil;
		case "enum":
			return "0", nil;
		case "string":
			return "", fmt.Errorf ("it is not possible for a struct to contain a string value");
		case "class", "optionalclass":
			return "", fmt.Errorf ("it is not possible for a struct to contain a handle value");
		default:
			return "", fmt.Errorf ("unknown member type %s", memberType);
	}

}


// CreateCAbiHeader creates a C header file for the component's API
func CreateCAbiHeader(component ComponentDefinition, CHeaderName string) (error) {
	hfile, err := CreateLanguageFile(CHeaderName, "  ");
	if (err != nil) {
		return err;
	}
	hfile.WriteCLicenseHeader (component,
		fmt.Sprintf ("This is an autogenerated plain C Header file in order to allow an easy\n use of %s", component.LibraryName),
		true);
	err = buildCAbiHeader(component, hfile, component.NameSpace, component.BaseName, false);
	return err;
}

func writeClassMethodsIntoCCPPHeader(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter, NameSpace string, useCPPTypes bool) (error) {
	w.Writeln("");
	w.Writeln("/*************************************************************************************************************************");
	w.Writeln(" Class definition for %s", class.ClassName);
	w.Writeln("**************************************************************************************************************************/");

	for j := 0; j < len(class.Methods); j++ {
		method := class.Methods[j];
		err := WriteCCPPAbiMethod (method, w, NameSpace, class.ClassName, false, false, useCPPTypes);
		if (err != nil) {
			return err;
		}
	}
	return nil
}

func buildCAbiHeader(component ComponentDefinition, w LanguageWriter, NameSpace string, BaseName string, useCPPTypes bool) (error) {
	sIncludeGuard :=  "__"+strings.ToUpper(NameSpace)+"_HEADER"
	if (useCPPTypes) {
		sIncludeGuard += "_CPP"
	}

	w.Writeln("#ifndef %s", sIncludeGuard);
	w.Writeln("#define %s", sIncludeGuard);
	w.Writeln("");

	w.Writeln("#ifdef __%s_EXPORTS", strings.ToUpper (NameSpace));
	w.Writeln("#ifdef _WIN32");
	w.Writeln("#define %s_DECLSPEC __declspec (dllexport)", strings.ToUpper (NameSpace));
	w.Writeln("#else // _WIN32");
	w.Writeln("#define %s_DECLSPEC __attribute__((visibility(\"default\")))", strings.ToUpper (NameSpace));
	w.Writeln("#endif // _WIN32");
	
	w.Writeln("#else // __%s_EXPORTS", strings.ToUpper (NameSpace));
	w.Writeln("#define %s_DECLSPEC", strings.ToUpper (NameSpace));
	w.Writeln("#endif // __%s_EXPORTS", strings.ToUpper (NameSpace));
	w.Writeln("");

	if (useCPPTypes) {
		w.Writeln("#include \"%s_types.hpp\"", BaseName);
	} else {
		w.Writeln("#include \"%s_types.h\"", BaseName);
	}
	w.Writeln("");
	for _, subComponent := range(component.ImportedComponentDefinitions) {
		w.Writeln("#include \"%s_dynamic.hpp\"", subComponent.BaseName)
	}
	w.Writeln("")


	w.Writeln("#ifdef __cplusplus");
	w.Writeln("extern \"C\" {");
	w.Writeln("#endif");

	for i := 0; i < len(component.Classes); i++ {
		class := component.Classes[i];
		err := writeClassMethodsIntoCCPPHeader(component, class, w, NameSpace, useCPPTypes)
		if (err != nil) {
			return err;
		}
	}

	w.Writeln("");
	w.Writeln("/*************************************************************************************************************************");
	w.Writeln(" Global functions");
	w.Writeln("**************************************************************************************************************************/");
	
	global := component.Global;
	for j := 0; j < len(global.Methods); j++ {
		method := global.Methods[j];
		err := WriteCCPPAbiMethod(method, w, NameSpace, "Wrapper", true, false, useCPPTypes);
		if (err != nil) {
			return err;
		}
	}
	
	w.Writeln("");
	w.Writeln("#ifdef __cplusplus");
	w.Writeln("}");
	w.Writeln("#endif");
	w.Writeln("");
	w.Writeln("#endif // %s", sIncludeGuard);
	w.Writeln("");
	
	return nil;
}


// GetCExportName How do we name the exports in the plain C DLL
func GetCExportName (NameSpace string, ClassName string, method ComponentDefinitionMethod, isGlobal bool) (string) {
	CMethodName := "";
	if isGlobal {
		CMethodName = fmt.Sprintf("%s_%s", strings.ToLower(NameSpace), strings.ToLower(method.MethodName))
	} else {
		CMethodName = fmt.Sprintf("%s_%s_%s", strings.ToLower(NameSpace), strings.ToLower(ClassName), strings.ToLower(method.MethodName))
	}
	
	return CMethodName;
}


// WriteCCPPAbiMethod writes an ABI method as a C-function
func WriteCCPPAbiMethod(method ComponentDefinitionMethod, w LanguageWriter, NameSpace string, ClassName string, isGlobal bool, writeCallbacks bool, useCPPTypes bool) (error) {
	CMethodName := "";
	CCallbackName := "";
	parameters := "";
	if (isGlobal) {
		CMethodName = fmt.Sprintf ("%s_%s", strings.ToLower (NameSpace), strings.ToLower (method.MethodName));
		CCallbackName = fmt.Sprintf ("P%s%sPtr", NameSpace, method.MethodName);
	} else {
		CMethodName = fmt.Sprintf ("%s_%s_%s", strings.ToLower (NameSpace), strings.ToLower (ClassName), strings.ToLower (method.MethodName));
		CCallbackName = fmt.Sprintf ("P%s%s_%sPtr", NameSpace, ClassName, method.MethodName);
		parameters = fmt.Sprintf ("%s_%s p%s", NameSpace, ClassName, ClassName);
	}

	w.Writeln("");
	w.Writeln("/**");
	w.Writeln("* %s", method.MethodDescription);
	w.Writeln("*");
	if (!isGlobal) {
		w.Writeln("* @param[in] p%s - %s instance.", ClassName, ClassName);
	}
	

	for k := 0; k < len(method.Params); k++ {
		param := method.Params [k];
		cParams, err := generateCCPPParameter(param, ClassName, method.MethodName, NameSpace, useCPPTypes);
		if (err != nil) {
			return err;
		}
		for _, cParam := range cParams {
			w.Writeln(cParam.ParamComment);
			if (parameters != "") {
				parameters = parameters + ", ";
			}
			parameters = parameters + cParam.ParamType + " " + cParam.ParamName;
		}
	}
	
	w.Writeln("* @return error code or 0 (success)");
	w.Writeln("*/");
	
	if (writeCallbacks) {
		w.Writeln("typedef %sResult (*%s) (%s);", NameSpace, CCallbackName, parameters);
	} else {
		w.Writeln("%s_DECLSPEC %sResult %s(%s);", strings.ToUpper(NameSpace), NameSpace, CMethodName, parameters);
	}
	
	return nil;
}

func buildCCPPStructs(component ComponentDefinition, w LanguageWriter, NameSpace string, useCPPTypes bool) (error) {
	if (len(component.Structs) == 0) {
		return nil
	}

	var err error
	
	w.Writeln("/*************************************************************************************************************************");
	w.Writeln(" Declaration of structs");
	w.Writeln("**************************************************************************************************************************/");
	w.Writeln("");
		
	w.Writeln("#pragma pack (1)");
	w.Writeln("");

	for i := 0; i < len(component.Structs); i++ {
		structinfo := component.Structs[i];
		if (useCPPTypes) {
			w.Writeln("typedef struct s%s {", structinfo.Name);
		} else {
			w.Writeln("typedef struct s%s%s {", NameSpace, structinfo.Name);
		}
		
		for j := 0; j < len(structinfo.Members); j++ {
			member := structinfo.Members[j];
			arraysuffix := "";
			if (member.Rows > 0) {
				if (member.Columns > 0) {
					arraysuffix = fmt.Sprintf ("[%d][%d]", member.Columns, member.Rows)
				} else {
					arraysuffix = fmt.Sprintf ("[%d]",member.Rows)
				}
			}
			var memberLine string
			if (useCPPTypes) {
				memberLine, err= getCPPMemberLine(member, NameSpace, arraysuffix, structinfo.Name, ";")
			} else {
				memberLine, err= getCMemberLine(member, NameSpace, arraysuffix, structinfo.Name)
			}
			if (err!=nil) {
				return err
			}
			w.Writeln("    %s", memberLine)
		}
		if (useCPPTypes) {
			w.Writeln("} s%s;", structinfo.Name);
		} else {
			w.Writeln("} s%s%s;", NameSpace, structinfo.Name);
		}
		w.Writeln("");
	}

	w.Writeln("#pragma pack ()");
	w.Writeln("");

	return nil
}

func buildCCPPEnums(component ComponentDefinition, w LanguageWriter, NameSpace string, useCPPTypes bool) (error) {
	if (len(component.Enums) == 0) {
		return nil
	}

	w.Writeln("/*************************************************************************************************************************");
	w.Writeln(" Declaration of enums");
	w.Writeln("**************************************************************************************************************************/");
	w.Writeln("");

	for i := 0; i < len(component.Enums); i++ {
		enum := component.Enums[i];
		if (useCPPTypes) {
			if (enum.Description != "") {
				w.Writeln("/**");
				w.Writeln("* enum class e%s - %s", enum.Name, enum.Description);
				w.Writeln("*/");
			}
			w.Writeln("enum class e%s : %s_int32 {", enum.Name, NameSpace);
		} else {
			if (enum.Description != "") {
				w.Writeln("/**");
				w.Writeln("* enum e%s%s - %s", NameSpace, enum.Name, enum.Description);
				w.Writeln("*/");
			}
			w.Writeln("typedef enum e%s%s {", NameSpace, enum.Name);
		}
		
		for j := 0; j < len(enum.Options); j++ {
			comma := "";
			if (j < len(enum.Options) - 1) {
				comma = ",";
			}
			option := enum.Options[j];
			if (useCPPTypes) {
				if (option.Description != "") {
					w.Writeln("  %s = %d%s /** %s */", option.Name, option.Value, comma, option.Description);
				} else {
					w.Writeln("  %s = %d%s", option.Name, option.Value, comma);
				}
			} else {
				if (option.Description != "") {
					w.Writeln("  e%s%s = %d%s /** %s */", enum.Name, option.Name, option.Value, comma, option.Description);
				} else {
					w.Writeln("  e%s%s = %d%s", enum.Name, option.Name, option.Value, comma);
				}
			}
		}
		if (useCPPTypes) {
			w.Writeln("};");
		} else {
			w.Writeln("} e%s%s;", NameSpace, enum.Name);
		}
		w.Writeln("");
	}
	
	if (!useCPPTypes) {
		w.Writeln("/*************************************************************************************************************************");
		w.Writeln(" Declaration of enum members for 4 byte struct alignment");
		w.Writeln("**************************************************************************************************************************/");
		w.Writeln("");

		for i := 0; i < len(component.Enums); i++ {
			enum := component.Enums[i];
			w.Writeln("typedef union {");
			w.Writeln("  e%s%s m_enum;", NameSpace, enum.Name);
			w.Writeln("  int m_code;");
			w.Writeln("} structEnum%s%s;", NameSpace, enum.Name);
			w.Writeln("");
		}
	}
	return nil
}


func buildCCPPFunctionPointers(component ComponentDefinition, w LanguageWriter, NameSpace string, useCPPTypes bool) (error) {
	if len(component.Functions) == 0 {
		return nil
	}

	w.Writeln("/*************************************************************************************************************************");
	w.Writeln(" Declaration of function pointers ");
	w.Writeln("**************************************************************************************************************************/");
	for i := 0; i < len(component.Functions); i++ {
		functiontype := component.Functions[i]
		returnType := "void"
		parameters := ""

		w.Writeln("");
		w.Writeln("/**");
		if (useCPPTypes) {
			w.Writeln("* %s - %s", functiontype.FunctionName, functiontype.FunctionDescription )
		} else {
			w.Writeln("* %s%s - %s", NameSpace, functiontype.FunctionName, functiontype.FunctionDescription )
		}
		
		w.Writeln("*")
		for j := 0; j < len(functiontype.Params); j++ {
			param := functiontype.Params[j]

			cParams, err := generateCCPPParameter(param, "", functiontype.FunctionName, NameSpace, useCPPTypes)
			if (err != nil) {
				return err;
			}
			for _, cParam := range cParams {
				w.Writeln(cParam.ParamComment);

				if (parameters != "") {
					parameters = parameters + ", "
				}
				parameters = parameters + cParam.ParamType
			}
		}
		w.Writeln("*/");
		if (useCPPTypes) {
			w.Writeln("typedef %s(*%s)(%s);", returnType, functiontype.FunctionName, parameters);
		} else {
			w.Writeln("typedef %s(*%s%s)(%s);", returnType, NameSpace, functiontype.FunctionName, parameters);
		}
	}
	w.Writeln("");
	return nil
}

func getCParameterTypeName(ParamTypeName string, NameSpace string, ParamClass string)(string, error) {
	paramNameSpace, paramClassName, err := decomposeParamClassName(ParamClass)
	if (err != nil) {
		return "", err
	}
	if len(paramNameSpace) == 0 {
		paramNameSpace = NameSpace
	}


	cParamTypeName := "";
	switch (ParamTypeName) {
		case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "single", "double":
			cParamTypeName = fmt.Sprintf ("%s_%s", paramNameSpace, ParamTypeName);
		case "bool":
			cParamTypeName = "bool";
		case "pointer":
			cParamTypeName = fmt.Sprintf ("%s_pvoid", paramNameSpace);
		case "string":
			cParamTypeName = "char *";
		case "enum":
			cParamTypeName = fmt.Sprintf ("e%s%s", paramNameSpace, paramClassName);
		case "struct":
			cParamTypeName = fmt.Sprintf ("s%s%s *", paramNameSpace, paramClassName);
		case "basicarray":
			basicTypeName, err := getCParameterTypeName(paramClassName, paramNameSpace, "");
			if (err != nil) {
				return "", err;
			}
			cParamTypeName = fmt.Sprintf ("%s *", basicTypeName);
		case "structarray":
			cParamTypeName = fmt.Sprintf ("s%s%s *", paramNameSpace, paramClassName)
		case "class", "optionalclass":
			cParamTypeName = fmt.Sprintf ("%s_%s", paramNameSpace, paramClassName)
		case "functiontype":
			cParamTypeName = fmt.Sprintf ("%s%s", paramNameSpace, paramClassName)
		default:
			return "", fmt.Errorf ("invalid parameter type \"%s\" for C-parameter", ParamTypeName);
	}
	return cParamTypeName, nil;
}

// CParameter is a handy representation of a function parameter in C
type CParameter struct {
	ParamType string
	ParamName string
	ParamComment string
	ParamDocumentationLine string
}


func generateCCPPParameter(param ComponentDefinitionParam, className string, methodName string, NameSpace string, useCPPTypes bool) ([]CParameter, error) {
	cParams := make([]CParameter,1)
	var cParamTypeName string
	var err error
	if (useCPPTypes) {
		cParamTypeName, err = getCPPParameterTypeName(param.ParamType, NameSpace, param.ParamClass);
	} else {
		cParamTypeName, err = getCParameterTypeName(param.ParamType, NameSpace, param.ParamClass);
	}
	if (err != nil) {
		return nil, err;
	}

	switch (param.ParamPass) {
	case "in":
		switch (param.ParamType) {
			case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "n" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamDocumentationLine = fmt.Sprintf(":param %s: %s", cParams[0].ParamName, param.ParamDescription)

			case "bool":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "b" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamDocumentationLine = fmt.Sprintf(":param %s: %s", cParams[0].ParamName, param.ParamDescription)
				
			case "single":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "f" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamDocumentationLine = fmt.Sprintf(":param %s: %s", cParams[0].ParamName, param.ParamDescription)

			case "double":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "d" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamDocumentationLine = fmt.Sprintf(":param %s: %s", cParams[0].ParamName, param.ParamDescription)

			case "pointer":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "p" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamDocumentationLine = fmt.Sprintf(":param %s: %s", cParams[0].ParamName, param.ParamDescription)
				
			case "string":
				cParams[0].ParamType = "const " + cParamTypeName;
				cParams[0].ParamName = "p" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamDocumentationLine = fmt.Sprintf(":param %s: %s", cParams[0].ParamName, param.ParamDescription)

			case "enum":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "e" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamDocumentationLine = fmt.Sprintf(":param %s: %s", cParams[0].ParamName, param.ParamDescription)

			case "struct":
				cParams[0].ParamType = "const " + cParamTypeName;
				cParams[0].ParamName = "p" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamDocumentationLine = fmt.Sprintf(":param %s: %s", cParams[0].ParamName, param.ParamDescription)

			case "basicarray", "structarray":
				cParams = make([]CParameter,2)
				cParams[0].ParamType = fmt.Sprintf ("%s_uint64", NameSpace);
				cParams[0].ParamName = "n" + param.ParamName + "BufferSize";
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - Number of elements in buffer", cParams[0].ParamName);
				cParams[0].ParamDocumentationLine = fmt.Sprintf(":param %s: %s", cParams[0].ParamName, param.ParamDescription)

				cParams[1].ParamType = "const " + cParamTypeName;
				cParams[1].ParamName = "p" + param.ParamName + "Buffer";
				cParams[1].ParamComment = fmt.Sprintf("* @param[in] %s - %s buffer of %s", cParams[1].ParamName, param.ParamClass, param.ParamDescription);
				cParams[1].ParamDocumentationLine = fmt.Sprintf(":param %s: buffer of %s", cParams[0].ParamName, param.ParamDescription)

			case "class", "optionalclass":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "p" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamDocumentationLine = fmt.Sprintf(":param %s: %s", cParams[0].ParamName, param.ParamDescription)

			case "functiontype":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "p" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamDocumentationLine = fmt.Sprintf(":param %s: %s", cParams[0].ParamName, param.ParamDescription)

			default:
				return nil, fmt.Errorf ("invalid method parameter type \"%s\" for %s.%s (%s)", param.ParamType, className, methodName, param.ParamName);
		}
	
	case "out", "return":
	
		switch (param.ParamType) {
		
			case "uint8", "uint16", "uint32", "uint64",  "int8", "int16", "int32", "int64", "bool", "single", "double", "pointer", "enum":
				cParams[0].ParamType = cParamTypeName + " *";
				cParams[0].ParamName = "p" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[out] %s - %s", cParams[0].ParamName, param.ParamDescription);
				if ("out" == param.ParamPass) {
					cParams[0].ParamDocumentationLine = fmt.Sprintf(":param %s: %s", cParams[0].ParamName, param.ParamDescription)
				} else {
					cParams[0].ParamDocumentationLine = fmt.Sprintf(":return: %s", param.ParamDescription)
				}

			case "struct":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "p" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[out] %s - %s", cParams[0].ParamName, param.ParamDescription);
				if ("out" == param.ParamPass) {
					cParams[0].ParamDocumentationLine = fmt.Sprintf(":param %s: %s", cParams[0].ParamName, param.ParamDescription)
				} else {
					cParams[0].ParamDocumentationLine = fmt.Sprintf(":return: %s", param.ParamDescription)
				}
				
			case "basicarray", "structarray":
				cParams = make([]CParameter,3)
				cParams[0].ParamType = fmt.Sprintf("const %s_uint64", NameSpace)
				cParams[0].ParamName = "n" + param.ParamName + "BufferSize";
				paramComment0 := "Number of elements in buffer"
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, paramComment0);

				cParams[1].ParamType = fmt.Sprintf("%s_uint64*", NameSpace)
				cParams[1].ParamName = "p" + param.ParamName + "NeededCount";
				paramComment1 := "will be filled with the count of the written elements, or needed buffer size."
				cParams[1].ParamComment = fmt.Sprintf("* @param[out] %s - %s", cParams[1].ParamName, paramComment1);

				cParams[2].ParamType = cParamTypeName;
				cParams[2].ParamName = "p" + param.ParamName + "Buffer";
				paramComment2 := "buffer of " + param.ParamDescription
				cParams[2].ParamComment = fmt.Sprintf("* @param[out] %s - %s  %s", cParams[2].ParamName, param.ParamClass, paramComment2);

				cParams[0].ParamDocumentationLine = fmt.Sprintf(":param %s: %s", cParams[0].ParamName, paramComment0)
				cParams[1].ParamDocumentationLine = fmt.Sprintf(":param %s: %s", cParams[1].ParamName, paramComment1)
				cParams[2].ParamDocumentationLine = fmt.Sprintf(":param %s: %s", cParams[2].ParamName, paramComment2)

			case "string":
				cParams = make([]CParameter,3)
				cParams[0].ParamType = fmt.Sprintf("const %s_uint32", NameSpace)
				cParams[0].ParamName = "n" + param.ParamName + "BufferSize";
				paramComment0 := "size of the buffer (including trailing 0)"
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, paramComment0);

				cParams[1].ParamType = fmt.Sprintf("%s_uint32*", NameSpace)
				cParams[1].ParamName = "p" + param.ParamName + "NeededChars";
				paramComment1 := "will be filled with the count of the written bytes, or needed buffer size."
				cParams[1].ParamComment = fmt.Sprintf("* @param[out] %s - %s", cParams[1].ParamName, paramComment1);

				cParams[2].ParamType = cParamTypeName;
				cParams[2].ParamName = "p" + param.ParamName + "Buffer";
				paramComment2 := fmt.Sprintf("buffer of %s, may be NULL", param.ParamDescription)
				cParams[2].ParamComment = fmt.Sprintf("* @param[out] %s - %s %s", cParams[2].ParamName, param.ParamClass, paramComment2);

				cParams[0].ParamDocumentationLine = fmt.Sprintf(":param %s: %s", cParams[0].ParamName, paramComment0)
				cParams[1].ParamDocumentationLine = fmt.Sprintf(":param %s: %s", cParams[1].ParamName, paramComment1)
				cParams[2].ParamDocumentationLine = fmt.Sprintf(":param %s: %s", cParams[2].ParamName, paramComment2)

			case "class", "optionalclass":
				cParams[0].ParamType = cParamTypeName + " *";
				cParams[0].ParamName = "p" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[out] %s - %s", cParams[0].ParamName, param.ParamDescription);
	
				if ("out" == param.ParamPass) {
					cParams[0].ParamDocumentationLine = fmt.Sprintf(":param %s: %s", cParams[0].ParamName, param.ParamDescription)
				} else {
					cParams[0].ParamDocumentationLine = fmt.Sprintf(":return: %s", param.ParamDescription)
				}

			default:
				return nil, fmt.Errorf ("invalid method parameter type \"%s\" for %s.%s (%s)", param.ParamType, className, methodName, param.ParamName);
		}
		
	default:
		return nil, fmt.Errorf ("invalid method parameter passing \"%s\" for %s.%s (%s)", param.ParamPass, className, methodName, param.ParamName);
	}

	return cParams, nil;
}

// GenerateCParameters generates an array of cParameters for a method
func GenerateCParameters(method ComponentDefinitionMethod, className string, NameSpace string) ([]CParameter, error) {
	parameters := []CParameter{}
	for k := 0; k < len(method.Params); k++ {
		param := method.Params [k]
		
		cParam, err := generateCCPPParameter(param, className, method.MethodName, NameSpace, false);
		if err != nil {
			return nil, err
		}
		parameters = append(parameters, cParam...)
	}

	return parameters, nil;
}
