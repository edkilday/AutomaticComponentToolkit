/*++

Copyright (C) 2021 ADSK

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.7.0-develop.

Abstract: This is an autogenerated Java file in order to allow an easy
 use of RTTI

Interface version: 1.0.0

*/

package rtti;

import com.sun.jna.*;

import java.nio.charset.StandardCharsets;


public class RTTIWrapper {

	public static class EnumConversion {
	}

	protected Function rtti_getversion;
	protected Function rtti_getlasterror;
	protected Function rtti_releaseinstance;
	protected Function rtti_acquireinstance;
	protected Function rtti_injectcomponent;
	protected Function rtti_getsymbollookupmethod;
	protected Function rtti_createzoo;
	protected Function rtti_base_classtypeid;
	protected Function rtti_animal_name;
	protected Function rtti_tiger_roar;
	protected Function rtti_animaliterator_getnextanimal;
	protected Function rtti_zoo_iterator;

	protected NativeLibrary mLibrary;

	public RTTIWrapper(String libraryPath) {
		mLibrary = NativeLibrary.getInstance(libraryPath);
		rtti_getversion = mLibrary.getFunction("rtti_getversion");
		rtti_getlasterror = mLibrary.getFunction("rtti_getlasterror");
		rtti_releaseinstance = mLibrary.getFunction("rtti_releaseinstance");
		rtti_acquireinstance = mLibrary.getFunction("rtti_acquireinstance");
		rtti_injectcomponent = mLibrary.getFunction("rtti_injectcomponent");
		rtti_getsymbollookupmethod = mLibrary.getFunction("rtti_getsymbollookupmethod");
		rtti_createzoo = mLibrary.getFunction("rtti_createzoo");
		rtti_base_classtypeid = mLibrary.getFunction("rtti_base_classtypeid");
		rtti_animal_name = mLibrary.getFunction("rtti_animal_name");
		rtti_tiger_roar = mLibrary.getFunction("rtti_tiger_roar");
		rtti_animaliterator_getnextanimal = mLibrary.getFunction("rtti_animaliterator_getnextanimal");
		rtti_zoo_iterator = mLibrary.getFunction("rtti_zoo_iterator");
	}

	public RTTIWrapper(Pointer lookupPointer) throws RTTIException {
		Function lookupMethod = Function.getFunction(lookupPointer);
		rtti_getversion = loadFunctionByLookup(lookupMethod, "rtti_getversion");
		rtti_getlasterror = loadFunctionByLookup(lookupMethod, "rtti_getlasterror");
		rtti_releaseinstance = loadFunctionByLookup(lookupMethod, "rtti_releaseinstance");
		rtti_acquireinstance = loadFunctionByLookup(lookupMethod, "rtti_acquireinstance");
		rtti_injectcomponent = loadFunctionByLookup(lookupMethod, "rtti_injectcomponent");
		rtti_getsymbollookupmethod = loadFunctionByLookup(lookupMethod, "rtti_getsymbollookupmethod");
		rtti_createzoo = loadFunctionByLookup(lookupMethod, "rtti_createzoo");
		rtti_base_classtypeid = loadFunctionByLookup(lookupMethod, "rtti_base_classtypeid");
		rtti_animal_name = loadFunctionByLookup(lookupMethod, "rtti_animal_name");
		rtti_tiger_roar = loadFunctionByLookup(lookupMethod, "rtti_tiger_roar");
		rtti_animaliterator_getnextanimal = loadFunctionByLookup(lookupMethod, "rtti_animaliterator_getnextanimal");
		rtti_zoo_iterator = loadFunctionByLookup(lookupMethod, "rtti_zoo_iterator");
	}

	protected void checkError(Base instance, int errorCode) throws RTTIException {
		if (instance != null && instance.mWrapper != this) {
			throw new RTTIException(RTTIException.RTTI_ERROR_INVALIDCAST, "invalid wrapper call");
		}
		if (errorCode != RTTIException.RTTI_SUCCESS) {
			if (instance != null) {
				GetLastErrorResult result = getLastError(instance);
				throw new RTTIException(errorCode, result.ErrorMessage);
			} else {
				throw new RTTIException(errorCode, "");
			}
		}
	}

	private Function loadFunctionByLookup(Function lookupMethod, String functionName) throws RTTIException {
		byte[] bytes = functionName.getBytes(StandardCharsets.UTF_8);
		Memory name = new Memory(bytes.length+1);
		name.write(0, bytes, 0, bytes.length);
		name.setByte(bytes.length, (byte)0);
		Pointer address = new Memory(8);
		java.lang.Object[] addressParam = new java.lang.Object[]{name, address};
		checkError(null, lookupMethod.invokeInt(addressParam));
		return Function.getFunction(address.getPointer(0));
	}

	/**
	 * retrieves the binary version of this library.
	 *
	 * @return GetVersion Result Tuple
	 * @throws RTTIException
	 */
	public GetVersionResult getVersion() throws RTTIException {
		Pointer bufferMajor = new Memory(4);
		Pointer bufferMinor = new Memory(4);
		Pointer bufferMicro = new Memory(4);
		checkError(null, rtti_getversion.invokeInt(new java.lang.Object[]{bufferMajor, bufferMinor, bufferMicro}));
		GetVersionResult returnTuple = new GetVersionResult();
		returnTuple.Major = bufferMajor.getInt(0);
		returnTuple.Minor = bufferMinor.getInt(0);
		returnTuple.Micro = bufferMicro.getInt(0);
		return returnTuple;
	}

	public static class GetVersionResult {
		/**
		 * returns the major version of this library
		 */
		public int Major;

		/**
		 * returns the minor version of this library
		 */
		public int Minor;

		/**
		 * returns the micro version of this library
		 */
		public int Micro;

	}
	/**
	 * Returns the last error recorded on this object
	 *
	 * @param instance Instance Handle
	 * @return GetLastError Result Tuple
	 * @throws RTTIException
	 */
	public GetLastErrorResult getLastError(Base instance) throws RTTIException {
		RTTIHandle.ByValue instanceHandle;
		if (instance != null) {
			instanceHandle = instance.getHandle().Value();
		} else {
			throw new RTTIException(RTTIException.RTTI_ERROR_INVALIDPARAM, "Instance is a null value.");
		}
		Pointer bytesNeededErrorMessage = new Memory(4);
		Pointer bufferHasError = new Memory(1);
		checkError(null, rtti_getlasterror.invokeInt(new java.lang.Object[]{instanceHandle, 0, bytesNeededErrorMessage, null, bufferHasError}));
		int sizeErrorMessage = bytesNeededErrorMessage.getInt(0);
		Pointer bufferErrorMessage = new Memory(sizeErrorMessage);
		checkError(null, rtti_getlasterror.invokeInt(new java.lang.Object[]{instanceHandle, sizeErrorMessage, bytesNeededErrorMessage, bufferErrorMessage, bufferHasError}));
		GetLastErrorResult returnTuple = new GetLastErrorResult();
		returnTuple.ErrorMessage = new String(bufferErrorMessage.getByteArray(0, sizeErrorMessage - 1), StandardCharsets.UTF_8);
		returnTuple.HasError = bufferHasError.getByte(0) != 0;
		return returnTuple;
	}

	public static class GetLastErrorResult {
		/**
		 * Message of the last error
		 */
		public String ErrorMessage;

		/**
		 * Is there a last error to query
		 */
		public boolean HasError;

	}
	/**
	 * Releases shared ownership of an Instance
	 *
	 * @param instance Instance Handle
	 * @throws RTTIException
	 */
	public void releaseInstance(Base instance) throws RTTIException {
		RTTIHandle.ByValue instanceHandle;
		if (instance != null) {
			instanceHandle = instance.getHandle().Value();
		} else {
			throw new RTTIException(RTTIException.RTTI_ERROR_INVALIDPARAM, "Instance is a null value.");
		}
		checkError(null, rtti_releaseinstance.invokeInt(new java.lang.Object[]{instanceHandle}));
	}

	/**
	 * Acquires shared ownership of an Instance
	 *
	 * @param instance Instance Handle
	 * @throws RTTIException
	 */
	public void acquireInstance(Base instance) throws RTTIException {
		RTTIHandle.ByValue instanceHandle;
		if (instance != null) {
			instanceHandle = instance.getHandle().Value();
		} else {
			throw new RTTIException(RTTIException.RTTI_ERROR_INVALIDPARAM, "Instance is a null value.");
		}
		checkError(null, rtti_acquireinstance.invokeInt(new java.lang.Object[]{instanceHandle}));
	}

	/**
	 * Injects an imported component for usage within this component
	 *
	 * @param nameSpace NameSpace of the injected component
	 * @param symbolAddressMethod Address of the SymbolAddressMethod of the injected component
	 * @throws RTTIException
	 */
	public void injectComponent(String nameSpace, Pointer symbolAddressMethod) throws RTTIException {
		byte[] bytesNameSpace = nameSpace.getBytes(StandardCharsets.UTF_8);
		Memory bufferNameSpace = new Memory(bytesNameSpace.length + 1);
		bufferNameSpace.write(0, bytesNameSpace, 0, bytesNameSpace.length);
		bufferNameSpace.setByte(bytesNameSpace.length, (byte)0);
		checkError(null, rtti_injectcomponent.invokeInt(new java.lang.Object[]{bufferNameSpace, symbolAddressMethod}));

		boolean nameSpaceFound = false;
		if (!nameSpaceFound) {
			throw new RTTIException(RTTIException.RTTI_ERROR_COULDNOTLOADLIBRARY, "Unknown namespace " + nameSpace);
		}
	}

	/**
	 * Returns the address of the SymbolLookupMethod
	 *
	 * @return Address of the SymbolAddressMethod
	 * @throws RTTIException
	 */
	public Pointer getSymbolLookupMethod() throws RTTIException {
		Pointer bufferSymbolLookupMethod = new Memory(8);
		checkError(null, rtti_getsymbollookupmethod.invokeInt(new java.lang.Object[]{bufferSymbolLookupMethod}));
		return bufferSymbolLookupMethod.getPointer(0);
	}

	/**
	 * Create a new zoo with animals
	 *
	 * @return 
	 * @throws RTTIException
	 */
	public Zoo createZoo() throws RTTIException {
		RTTIHandle handleInstance = new RTTIHandle();
		checkError(null, rtti_createzoo.invokeInt(new java.lang.Object[]{handleInstance}));
		Zoo instance = null;
		if (handleInstance.Handle == Pointer.NULL) {
		  throw new RTTIException(RTTIException.RTTI_ERROR_INVALIDPARAM, "Instance was a null pointer");
		}
		instance = this.PolymorphicFactory(handleInstance, Zoo.class);
		return instance;
	}

	public <T> T PolymorphicFactory(RTTIHandle handle, Class<T> cls) {
		Class[] cArg = new Class[2];
		cArg[0] = RTTIWrapper.class;
		cArg[1] = RTTIHandle.class;
	
		try {
			T obj = null;
			int msbId = (int)(handle.ClassTypeId >> 32); 
			int lsbId = (int)handle.ClassTypeId; 
			switch(msbId) {
				case 0x1549AD28: 
					switch(lsbId) {
						case 0x813DAE05: obj = (T)(new Base(this, handle)); break; // First 64 bits of SHA1 of a string: "RTTI::Base"
					}
				break;
				case 0xF1917FE6: 
					switch(lsbId) {
						case 0xBBE77831: obj = (T)(new AnimalIterator(this, handle)); break; // First 64 bits of SHA1 of a string: "RTTI::AnimalIterator"
					}
				break;
				case 0x2262ABE8: 
					switch(lsbId) {
						case 0x0A5E7878: obj = (T)(new Zoo(this, handle)); break; // First 64 bits of SHA1 of a string: "RTTI::Zoo"
					}
				break;
				case 0x8B40467D: 
					switch(lsbId) {
						case 0xA6D327AF: obj = (T)(new Animal(this, handle)); break; // First 64 bits of SHA1 of a string: "RTTI::Animal"
					}
				break;
				case 0xBC9D5FA7: 
					switch(lsbId) {
						case 0x750C1020: obj = (T)(new Mammal(this, handle)); break; // First 64 bits of SHA1 of a string: "RTTI::Mammal"
					}
				break;
				case 0x6756AA8E: 
					switch(lsbId) {
						case 0xA5802EC3: obj = (T)(new Reptile(this, handle)); break; // First 64 bits of SHA1 of a string: "RTTI::Reptile"
					}
				break;
				case 0x9751971B: 
					switch(lsbId) {
						case 0xD2C2D958: obj = (T)(new Giraffe(this, handle)); break; // First 64 bits of SHA1 of a string: "RTTI::Giraffe"
					}
				break;
				case 0x08D007E7: 
					switch(lsbId) {
						case 0xB5F7BAF4: obj = (T)(new Tiger(this, handle)); break; // First 64 bits of SHA1 of a string: "RTTI::Tiger"
					}
				break;
				case 0x5F6826EF: 
					switch(lsbId) {
						case 0x909803B2: obj = (T)(new Snake(this, handle)); break; // First 64 bits of SHA1 of a string: "RTTI::Snake"
					}
				break;
				case 0x8E551B20: 
					switch(lsbId) {
						case 0x8A2E8321: obj = (T)(new Turtle(this, handle)); break; // First 64 bits of SHA1 of a string: "RTTI::Turtle"
					}
				break;
			}
			return obj;
		}
		catch(Exception e) {
			return null;
		}
	}
}

