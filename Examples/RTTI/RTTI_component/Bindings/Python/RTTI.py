'''++

Copyright (C) 2020 ADSK

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.7.0-develop.

Abstract: This is an autogenerated Python file in order to allow an easy
 use of RTTI

Interface version: 1.0.0

'''


import ctypes
import platform
import enum
import os

name = "rtti"

class RTTIHandle(ctypes.Structure):
	''' creates a struct to match RTTIHandle C type '''
	
	_fields_ = [('Handle', ctypes.c_void_p),
		('ClassTypeId', ctypes.c_uint64)]

'''Definition of domain specific exception
'''
class ERTTIException(Exception):
	def __init__(self, code, message = ''):
		self._code = code
		self._message = message
	
	def __str__(self):
		if self._message:
			return 'RTTIException ' + str(self._code) + ': '+ str(self._message)
		return 'RTTIException ' + str(self._code)

'''Definition of binding API version
'''
class BindingVersion(enum.IntEnum):
	MAJOR = 1
	MINOR = 0
	MICRO = 0

'''Definition Error Codes
'''
class ErrorCodes(enum.IntEnum):
	SUCCESS = 0
	NOTIMPLEMENTED = 1
	INVALIDPARAM = 2
	INVALIDCAST = 3
	BUFFERTOOSMALL = 4
	GENERICEXCEPTION = 5
	COULDNOTLOADLIBRARY = 6
	COULDNOTFINDLIBRARYEXPORT = 7
	INCOMPATIBLEBINARYVERSION = 8

'''Definition of Function Table
'''
class FunctionTable:
	rtti_getversion = None
	rtti_getlasterror = None
	rtti_releaseinstance = None
	rtti_acquireinstance = None
	rtti_injectcomponent = None
	rtti_getsymbollookupmethod = None
	rtti_createzoo = None
	rtti_base_classtypeid = None
	rtti_animal_name = None
	rtti_tiger_roar = None
	rtti_animaliterator_getnextanimal = None
	rtti_zoo_iterator = None

'''Definition of Structs
'''
'''Definition of TestStruct
'''
class TestStruct(ctypes.Structure):
	_pack_ = 1
	_fields_ = [
		("X", ctypes.c_int32), 
		("Y", ctypes.c_int32)
	]


'''Wrapper Class Implementation
'''
class Wrapper:

	def __init__(self, libraryName = None, symbolLookupMethodAddress = None):
		ending = ''
		if platform.system() == 'Windows':
			ending = 'dll'
		elif platform.system() == 'Linux':
			ending = 'so'
		elif platform.system() == 'Darwin':
			ending = 'dylib'
		else:
			raise ERTTIException(ErrorCodes.COULDNOTLOADLIBRARY)
		
		if (not libraryName) and (not symbolLookupMethodAddress):
			libraryName = os.path.join(os.path.dirname(os.path.realpath(__file__)),'rtti')
		
		if libraryName is not None:
			path = libraryName + '.' + ending
			try:
				self.lib = ctypes.CDLL(path)
			except Exception as e:
				raise ERTTIException(ErrorCodes.COULDNOTLOADLIBRARY, str(e) + '| "'+path + '"' )
			
			self._loadFunctionTable()
		elif symbolLookupMethodAddress is not None:
				self.lib = FunctionTable()
				self._loadFunctionTableFromMethod(symbolLookupMethodAddress)
		else:
			raise ERTTIException(ErrorCodes.COULDNOTLOADLIBRARY, str(e))
		
		self._checkBinaryVersion()
	
	def _loadFunctionTableFromMethod(self, symbolLookupMethodAddress):
		try:
			symbolLookupMethodType = ctypes.CFUNCTYPE(ctypes.c_int32, ctypes.c_char_p, ctypes.POINTER(ctypes.c_void_p))
			symbolLookupMethod = symbolLookupMethodType(int(symbolLookupMethodAddress))
			
			methodAddress = ctypes.c_void_p()
			
			err = symbolLookupMethod(ctypes.c_char_p(str.encode("rtti_getversion")), methodAddress)
			if err != 0:
				raise ERTTIException(ErrorCodes.COULDNOTLOADLIBRARY, str(err))
			methodType = ctypes.CFUNCTYPE(ctypes.c_int32, ctypes.POINTER(ctypes.c_uint32), ctypes.POINTER(ctypes.c_uint32), ctypes.POINTER(ctypes.c_uint32))
			self.lib.rtti_getversion = methodType(int(methodAddress.value))
			
			err = symbolLookupMethod(ctypes.c_char_p(str.encode("rtti_getlasterror")), methodAddress)
			if err != 0:
				raise ERTTIException(ErrorCodes.COULDNOTLOADLIBRARY, str(err))
			methodType = ctypes.CFUNCTYPE(ctypes.c_int32, RTTIHandle, ctypes.c_uint64, ctypes.POINTER(ctypes.c_uint64), ctypes.c_char_p, ctypes.POINTER(ctypes.c_bool))
			self.lib.rtti_getlasterror = methodType(int(methodAddress.value))
			
			err = symbolLookupMethod(ctypes.c_char_p(str.encode("rtti_releaseinstance")), methodAddress)
			if err != 0:
				raise ERTTIException(ErrorCodes.COULDNOTLOADLIBRARY, str(err))
			methodType = ctypes.CFUNCTYPE(ctypes.c_int32, RTTIHandle)
			self.lib.rtti_releaseinstance = methodType(int(methodAddress.value))
			
			err = symbolLookupMethod(ctypes.c_char_p(str.encode("rtti_acquireinstance")), methodAddress)
			if err != 0:
				raise ERTTIException(ErrorCodes.COULDNOTLOADLIBRARY, str(err))
			methodType = ctypes.CFUNCTYPE(ctypes.c_int32, RTTIHandle)
			self.lib.rtti_acquireinstance = methodType(int(methodAddress.value))
			
			err = symbolLookupMethod(ctypes.c_char_p(str.encode("rtti_injectcomponent")), methodAddress)
			if err != 0:
				raise ERTTIException(ErrorCodes.COULDNOTLOADLIBRARY, str(err))
			methodType = ctypes.CFUNCTYPE(ctypes.c_int32, ctypes.c_char_p, ctypes.c_void_p)
			self.lib.rtti_injectcomponent = methodType(int(methodAddress.value))
			
			err = symbolLookupMethod(ctypes.c_char_p(str.encode("rtti_getsymbollookupmethod")), methodAddress)
			if err != 0:
				raise ERTTIException(ErrorCodes.COULDNOTLOADLIBRARY, str(err))
			methodType = ctypes.CFUNCTYPE(ctypes.c_int32, ctypes.POINTER(ctypes.c_void_p))
			self.lib.rtti_getsymbollookupmethod = methodType(int(methodAddress.value))
			
			err = symbolLookupMethod(ctypes.c_char_p(str.encode("rtti_createzoo")), methodAddress)
			if err != 0:
				raise ERTTIException(ErrorCodes.COULDNOTLOADLIBRARY, str(err))
			methodType = ctypes.CFUNCTYPE(ctypes.c_int32, ctypes.POINTER(RTTIHandle))
			self.lib.rtti_createzoo = methodType(int(methodAddress.value))
			
			err = symbolLookupMethod(ctypes.c_char_p(str.encode("rtti_base_classtypeid")), methodAddress)
			if err != 0:
				raise ERTTIException(ErrorCodes.COULDNOTLOADLIBRARY, str(err))
			methodType = ctypes.CFUNCTYPE(ctypes.c_int32, RTTIHandle, ctypes.POINTER(ctypes.c_uint64))
			self.lib.rtti_base_classtypeid = methodType(int(methodAddress.value))
			
			err = symbolLookupMethod(ctypes.c_char_p(str.encode("rtti_animal_name")), methodAddress)
			if err != 0:
				raise ERTTIException(ErrorCodes.COULDNOTLOADLIBRARY, str(err))
			methodType = ctypes.CFUNCTYPE(ctypes.c_int32, RTTIHandle, ctypes.c_uint64, ctypes.POINTER(ctypes.c_uint64), ctypes.c_char_p)
			self.lib.rtti_animal_name = methodType(int(methodAddress.value))
			
			err = symbolLookupMethod(ctypes.c_char_p(str.encode("rtti_tiger_roar")), methodAddress)
			if err != 0:
				raise ERTTIException(ErrorCodes.COULDNOTLOADLIBRARY, str(err))
			methodType = ctypes.CFUNCTYPE(ctypes.c_int32, RTTIHandle)
			self.lib.rtti_tiger_roar = methodType(int(methodAddress.value))
			
			err = symbolLookupMethod(ctypes.c_char_p(str.encode("rtti_animaliterator_getnextanimal")), methodAddress)
			if err != 0:
				raise ERTTIException(ErrorCodes.COULDNOTLOADLIBRARY, str(err))
			methodType = ctypes.CFUNCTYPE(ctypes.c_int32, RTTIHandle, ctypes.POINTER(RTTIHandle))
			self.lib.rtti_animaliterator_getnextanimal = methodType(int(methodAddress.value))
			
			err = symbolLookupMethod(ctypes.c_char_p(str.encode("rtti_zoo_iterator")), methodAddress)
			if err != 0:
				raise ERTTIException(ErrorCodes.COULDNOTLOADLIBRARY, str(err))
			methodType = ctypes.CFUNCTYPE(ctypes.c_int32, RTTIHandle, ctypes.POINTER(RTTIHandle))
			self.lib.rtti_zoo_iterator = methodType(int(methodAddress.value))
			
		except AttributeError as ae:
			raise ERTTIException(ErrorCodes.COULDNOTFINDLIBRARYEXPORT, ae.args[0])
		
	def _loadFunctionTable(self):
		try:
			self.lib.rtti_getversion.restype = ctypes.c_int32
			self.lib.rtti_getversion.argtypes = [ctypes.POINTER(ctypes.c_uint32), ctypes.POINTER(ctypes.c_uint32), ctypes.POINTER(ctypes.c_uint32)]
			
			self.lib.rtti_getlasterror.restype = ctypes.c_int32
			self.lib.rtti_getlasterror.argtypes = [RTTIHandle, ctypes.c_uint64, ctypes.POINTER(ctypes.c_uint64), ctypes.c_char_p, ctypes.POINTER(ctypes.c_bool)]
			
			self.lib.rtti_releaseinstance.restype = ctypes.c_int32
			self.lib.rtti_releaseinstance.argtypes = [RTTIHandle]
			
			self.lib.rtti_acquireinstance.restype = ctypes.c_int32
			self.lib.rtti_acquireinstance.argtypes = [RTTIHandle]
			
			self.lib.rtti_injectcomponent.restype = ctypes.c_int32
			self.lib.rtti_injectcomponent.argtypes = [ctypes.c_char_p, ctypes.c_void_p]
			
			self.lib.rtti_getsymbollookupmethod.restype = ctypes.c_int32
			self.lib.rtti_getsymbollookupmethod.argtypes = [ctypes.POINTER(ctypes.c_void_p)]
			
			self.lib.rtti_createzoo.restype = ctypes.c_int32
			self.lib.rtti_createzoo.argtypes = [ctypes.POINTER(RTTIHandle)]
			
			self.lib.rtti_base_classtypeid.restype = ctypes.c_int32
			self.lib.rtti_base_classtypeid.argtypes = [RTTIHandle, ctypes.POINTER(ctypes.c_uint64)]
			
			self.lib.rtti_animal_name.restype = ctypes.c_int32
			self.lib.rtti_animal_name.argtypes = [RTTIHandle, ctypes.c_uint64, ctypes.POINTER(ctypes.c_uint64), ctypes.c_char_p]
			
			self.lib.rtti_tiger_roar.restype = ctypes.c_int32
			self.lib.rtti_tiger_roar.argtypes = [RTTIHandle]
			
			self.lib.rtti_animaliterator_getnextanimal.restype = ctypes.c_int32
			self.lib.rtti_animaliterator_getnextanimal.argtypes = [RTTIHandle, ctypes.POINTER(RTTIHandle)]
			
			self.lib.rtti_zoo_iterator.restype = ctypes.c_int32
			self.lib.rtti_zoo_iterator.argtypes = [RTTIHandle, ctypes.POINTER(RTTIHandle)]
			
		except AttributeError as ae:
			raise ERTTIException(ErrorCodes.COULDNOTFINDLIBRARYEXPORT, ae.args[0])
	
	def _checkBinaryVersion(self):
		nMajor, nMinor, _ = self.GetVersion()
		if (nMajor != BindingVersion.MAJOR) or (nMinor < BindingVersion.MINOR):
			raise ERTTIException(ErrorCodes.INCOMPATIBLEBINARYVERSION)
	
	def checkError(self, instance, errorCode):
		if errorCode != ErrorCodes.SUCCESS.value:
			if instance:
				if instance._wrapper != self:
					raise ERTTIException(ErrorCodes.INVALIDCAST, 'invalid wrapper call')
			message,_ = self.GetLastError(instance)
			raise ERTTIException(errorCode, message)
	
	def GetVersion(self):
		pMajor = ctypes.c_uint32()
		pMinor = ctypes.c_uint32()
		pMicro = ctypes.c_uint32()
		self.checkError(None, self.lib.rtti_getversion(pMajor, pMinor, pMicro))
		
		return pMajor.value, pMinor.value, pMicro.value
	
	def GetLastError(self, InstanceObject):
		InstanceHandle = None
		if InstanceObject:
			InstanceHandle = InstanceObject._handle
		else:
			raise ERTTIException(ErrorCodes.INVALIDPARAM, 'Invalid return/output value')
		nErrorMessageBufferSize = ctypes.c_uint64(0)
		nErrorMessageNeededChars = ctypes.c_uint64(0)
		pErrorMessageBuffer = ctypes.c_char_p(None)
		pHasError = ctypes.c_bool()
		self.checkError(None, self.lib.rtti_getlasterror(InstanceHandle, nErrorMessageBufferSize, nErrorMessageNeededChars, pErrorMessageBuffer, pHasError))
		nErrorMessageBufferSize = ctypes.c_uint64(nErrorMessageNeededChars.value)
		pErrorMessageBuffer = (ctypes.c_char * (nErrorMessageNeededChars.value))()
		self.checkError(None, self.lib.rtti_getlasterror(InstanceHandle, nErrorMessageBufferSize, nErrorMessageNeededChars, pErrorMessageBuffer, pHasError))
		
		return pErrorMessageBuffer.value.decode(), pHasError.value
	
	def ReleaseInstance(self, InstanceObject):
		InstanceHandle = None
		if InstanceObject:
			InstanceHandle = InstanceObject._handle
		else:
			raise ERTTIException(ErrorCodes.INVALIDPARAM, 'Invalid return/output value')
		self.checkError(None, self.lib.rtti_releaseinstance(InstanceHandle))
		
	
	def AcquireInstance(self, InstanceObject):
		InstanceHandle = None
		if InstanceObject:
			InstanceHandle = InstanceObject._handle
		else:
			raise ERTTIException(ErrorCodes.INVALIDPARAM, 'Invalid return/output value')
		self.checkError(None, self.lib.rtti_acquireinstance(InstanceHandle))
		
	
	def InjectComponent(self, NameSpace, SymbolAddressMethod):
		pNameSpace = ctypes.c_char_p(str.encode(NameSpace))
		pSymbolAddressMethod = ctypes.c_void_p(SymbolAddressMethod)
		self.checkError(None, self.lib.rtti_injectcomponent(pNameSpace, pSymbolAddressMethod))
		
		bNameSpaceFound = False
		if not bNameSpaceFound:
			raise ERTTIException(ErrorCodes.COULDNOTLOADLIBRARY, "Unknown namespace " + NameSpace)
		
	
	def GetSymbolLookupMethod(self):
		pSymbolLookupMethod = ctypes.c_void_p()
		self.checkError(None, self.lib.rtti_getsymbollookupmethod(pSymbolLookupMethod))
		
		return pSymbolLookupMethod.value
	
	def CreateZoo(self):
		InstanceHandle = RTTIHandle()
		self.checkError(None, self.lib.rtti_createzoo(InstanceHandle))
		if InstanceHandle.Handle:
			InstanceObject = self._polymorphicFactory(InstanceHandle)
		else:
			raise ERTTIException(ErrorCodes.INVALIDCAST, 'Invalid return/output value')
		
		return InstanceObject
	
	def _polymorphicFactory(self, handle):
		class PolymorphicFactory():
			def getObjectById(self, handle, wrapper):
				methodName = 'getObjectById_' + format(handle.ClassTypeId, '016X')
				method = getattr(self, methodName, lambda: 'Invalid class type id')
				return method(handle, wrapper)
			def getObjectById_1549AD28813DAE05(self, handle, wrapper): # First 64 bits of SHA1 of a string: "RTTI::Base"
				return Base(handle, wrapper)
			def getObjectById_8B40467DA6D327AF(self, handle, wrapper): # First 64 bits of SHA1 of a string: "RTTI::Animal"
				return Animal(handle, wrapper)
			def getObjectById_BC9D5FA7750C1020(self, handle, wrapper): # First 64 bits of SHA1 of a string: "RTTI::Mammal"
				return Mammal(handle, wrapper)
			def getObjectById_6756AA8EA5802EC3(self, handle, wrapper): # First 64 bits of SHA1 of a string: "RTTI::Reptile"
				return Reptile(handle, wrapper)
			def getObjectById_9751971BD2C2D958(self, handle, wrapper): # First 64 bits of SHA1 of a string: "RTTI::Giraffe"
				return Giraffe(handle, wrapper)
			def getObjectById_08D007E7B5F7BAF4(self, handle, wrapper): # First 64 bits of SHA1 of a string: "RTTI::Tiger"
				return Tiger(handle, wrapper)
			def getObjectById_5F6826EF909803B2(self, handle, wrapper): # First 64 bits of SHA1 of a string: "RTTI::Snake"
				return Snake(handle, wrapper)
			def getObjectById_8E551B208A2E8321(self, handle, wrapper): # First 64 bits of SHA1 of a string: "RTTI::Turtle"
				return Turtle(handle, wrapper)
			def getObjectById_F1917FE6BBE77831(self, handle, wrapper): # First 64 bits of SHA1 of a string: "RTTI::AnimalIterator"
				return AnimalIterator(handle, wrapper)
			def getObjectById_2262ABE80A5E7878(self, handle, wrapper): # First 64 bits of SHA1 of a string: "RTTI::Zoo"
				return Zoo(handle, wrapper)
		
		factory = PolymorphicFactory()
		return factory.getObjectById(handle, self)
	


''' Class Implementation for Base
'''
class Base:
	def __init__(self, handle, wrapper):
		if not handle or not wrapper:
			raise ERTTIException(ErrorCodes.INVALIDPARAM)
		self._handle = handle
		self._wrapper = wrapper
	
	def __del__(self):
		self._wrapper.ReleaseInstance(self)
	def ClassTypeId(self):
		pClassTypeId = ctypes.c_uint64()
		self._wrapper.checkError(self, self._wrapper.lib.rtti_base_classtypeid(self._handle, pClassTypeId))
		
		return pClassTypeId.value
	


''' Class Implementation for Animal
'''
class Animal(Base):
	def __init__(self, handle, wrapper):
		Base.__init__(self, handle, wrapper)
	def Name(self):
		nResultBufferSize = ctypes.c_uint64(0)
		nResultNeededChars = ctypes.c_uint64(0)
		pResultBuffer = ctypes.c_char_p(None)
		self._wrapper.checkError(self, self._wrapper.lib.rtti_animal_name(self._handle, nResultBufferSize, nResultNeededChars, pResultBuffer))
		nResultBufferSize = ctypes.c_uint64(nResultNeededChars.value)
		pResultBuffer = (ctypes.c_char * (nResultNeededChars.value))()
		self._wrapper.checkError(self, self._wrapper.lib.rtti_animal_name(self._handle, nResultBufferSize, nResultNeededChars, pResultBuffer))
		
		return pResultBuffer.value.decode()
	


''' Class Implementation for Mammal
'''
class Mammal(Animal):
	def __init__(self, handle, wrapper):
		Animal.__init__(self, handle, wrapper)


''' Class Implementation for Reptile
'''
class Reptile(Animal):
	def __init__(self, handle, wrapper):
		Animal.__init__(self, handle, wrapper)


''' Class Implementation for Giraffe
'''
class Giraffe(Mammal):
	def __init__(self, handle, wrapper):
		Mammal.__init__(self, handle, wrapper)


''' Class Implementation for Tiger
'''
class Tiger(Mammal):
	def __init__(self, handle, wrapper):
		Mammal.__init__(self, handle, wrapper)
	def Roar(self):
		self._wrapper.checkError(self, self._wrapper.lib.rtti_tiger_roar(self._handle))
		
	


''' Class Implementation for Snake
'''
class Snake(Reptile):
	def __init__(self, handle, wrapper):
		Reptile.__init__(self, handle, wrapper)


''' Class Implementation for Turtle
'''
class Turtle(Reptile):
	def __init__(self, handle, wrapper):
		Reptile.__init__(self, handle, wrapper)


''' Class Implementation for AnimalIterator
'''
class AnimalIterator(Base):
	def __init__(self, handle, wrapper):
		Base.__init__(self, handle, wrapper)
	def GetNextAnimal(self):
		AnimalHandle = RTTIHandle()
		self._wrapper.checkError(self, self._wrapper.lib.rtti_animaliterator_getnextanimal(self._handle, AnimalHandle))
		if AnimalHandle.Handle:
			AnimalObject = self._wrapper._polymorphicFactory(AnimalHandle)
		else:
			AnimalObject = None
		
		return AnimalObject
	


''' Class Implementation for Zoo
'''
class Zoo(Base):
	def __init__(self, handle, wrapper):
		Base.__init__(self, handle, wrapper)
	def Iterator(self):
		IteratorHandle = RTTIHandle()
		self._wrapper.checkError(self, self._wrapper.lib.rtti_zoo_iterator(self._handle, IteratorHandle))
		if IteratorHandle.Handle:
			IteratorObject = self._wrapper._polymorphicFactory(IteratorHandle)
		else:
			raise ERTTIException(ErrorCodes.INVALIDCAST, 'Invalid return/output value')
		
		return IteratorObject
	
		
