// +build windows

package visa

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	visa32                   = syscall.NewLazyDLL("visa32.dll")
	procViAssertIntrSignal   = visa32.NewProc("viAssertIntrSignal")
	procViAssertTrigger      = visa32.NewProc("viAssertTrigger")
	procViAssertUtilSignal   = visa32.NewProc("viAssertUtilSignal")
	procViBufRead            = visa32.NewProc("viBufRead")
	procViBufWrite           = visa32.NewProc("viBufWrite")
	procViClear              = visa32.NewProc("viClear")
	procViClose              = visa32.NewProc("viClose")
	procViDisableEvent       = visa32.NewProc("viDisableEvent")
	procViDiscardEvents      = visa32.NewProc("viDiscardEvents")
	procViEnableEvent        = visa32.NewProc("viEnableEvent")
	procViFindNext           = visa32.NewProc("viFindNext")
	procViFindRsrc           = visa32.NewProc("viFindRsrc")
	procViFlush              = visa32.NewProc("viFlush")
	procViGetAttribute       = visa32.NewProc("viGetAttribute")
	procViGetDefaultRM       = visa32.NewProc("viGetDefaultRM")
	procViGpibCommand        = visa32.NewProc("viGpibCommand")
	procViGpibControlATN     = visa32.NewProc("viGpibControlATN")
	procViGpibControlREN     = visa32.NewProc("viGpibControlREN")
	procViGpibPassControl    = visa32.NewProc("viGpibPassControl")
	procViGpibSendIFC        = visa32.NewProc("viGpibSendIFC")
	procViIn16               = visa32.NewProc("viIn16")
	procViIn16Ex             = visa32.NewProc("viIn16Ex")
	procViIn32               = visa32.NewProc("viIn32")
	procViIn32Ex             = visa32.NewProc("viIn32Ex")
	procViIn64               = visa32.NewProc("viIn64")
	procViIn64Ex             = visa32.NewProc("viIn64Ex")
	procViIn8                = visa32.NewProc("viIn8")
	procViIn8Ex              = visa32.NewProc("viIn8Ex")
	procViInstallHandler     = visa32.NewProc("viInstallHandler")
	procViLock               = visa32.NewProc("viLock")
	procViMapAddress         = visa32.NewProc("viMapAddress")
	procViMapAddressEx       = visa32.NewProc("viMapAddressEx")
	procViMapTrigger         = visa32.NewProc("viMapTrigger")
	procViMemAlloc           = visa32.NewProc("viMemAlloc")
	procViMemAllocEx         = visa32.NewProc("viMemAllocEx")
	procViMemFree            = visa32.NewProc("viMemFree")
	procViMemFreeEx          = visa32.NewProc("viMemFreeEx")
	procViMove               = visa32.NewProc("viMove")
	procViMoveAsync          = visa32.NewProc("viMoveAsync")
	procViMoveAsyncEx        = visa32.NewProc("viMoveAsyncEx")
	procViMoveEx             = visa32.NewProc("viMoveEx")
	procViMoveIn16           = visa32.NewProc("viMoveIn16")
	procViMoveIn16Ex         = visa32.NewProc("viMoveIn16Ex")
	procViMoveIn32           = visa32.NewProc("viMoveIn32")
	procViMoveIn32Ex         = visa32.NewProc("viMoveIn32Ex")
	procViMoveIn64           = visa32.NewProc("viMoveIn64")
	procViMoveIn64Ex         = visa32.NewProc("viMoveIn64Ex")
	procViMoveIn8            = visa32.NewProc("viMoveIn8")
	procViMoveIn8Ex          = visa32.NewProc("viMoveIn8Ex")
	procViMoveOut16          = visa32.NewProc("viMoveOut16")
	procViMoveOut16Ex        = visa32.NewProc("viMoveOut16Ex")
	procViMoveOut32          = visa32.NewProc("viMoveOut32")
	procViMoveOut32Ex        = visa32.NewProc("viMoveOut32Ex")
	procViMoveOut64          = visa32.NewProc("viMoveOut64")
	procViMoveOut64Ex        = visa32.NewProc("viMoveOut64Ex")
	procViMoveOut8           = visa32.NewProc("viMoveOut8")
	procViMoveOut8Ex         = visa32.NewProc("viMoveOut8Ex")
	procViOpen               = visa32.NewProc("viOpen")
	procViOpenDefaultRM      = visa32.NewProc("viOpenDefaultRM")
	procViOut16              = visa32.NewProc("viOut16")
	procViOut16Ex            = visa32.NewProc("viOut16Ex")
	procViOut32              = visa32.NewProc("viOut32")
	procViOut32Ex            = visa32.NewProc("viOut32Ex")
	procViOut64              = visa32.NewProc("viOut64")
	procViOut64Ex            = visa32.NewProc("viOut64Ex")
	procViOut8               = visa32.NewProc("viOut8")
	procViOut8Ex             = visa32.NewProc("viOut8Ex")
	procViParseRsrc          = visa32.NewProc("viParseRsrc")
	procViParseRsrcEx        = visa32.NewProc("viParseRsrcEx")
	procViPeek16             = visa32.NewProc("viPeek16")
	procViPeek32             = visa32.NewProc("viPeek32")
	procViPeek64             = visa32.NewProc("viPeek64")
	procViPeek8              = visa32.NewProc("viPeek8")
	procViPoke16             = visa32.NewProc("viPoke16")
	procViPoke32             = visa32.NewProc("viPoke32")
	procViPoke64             = visa32.NewProc("viPoke64")
	procViPoke8              = visa32.NewProc("viPoke8")
	procViPrintf             = visa32.NewProc("viPrintf")
	procViPxiReserveTriggers = visa32.NewProc("viPxiReserveTriggers")
	procViQueryf             = visa32.NewProc("viQueryf")
	procViRead               = visa32.NewProc("viRead")
	procViReadAsync          = visa32.NewProc("viReadAsync")
	procViReadSTB            = visa32.NewProc("viReadSTB")
	procViReadToFile         = visa32.NewProc("viReadToFile")
	procViSPrintf            = visa32.NewProc("viSPrintf")
	procViSScanf             = visa32.NewProc("viSScanf")
	procViScanf              = visa32.NewProc("viScanf")
	procViSetAttribute       = visa32.NewProc("viSetAttribute")
	procViSetBuf             = visa32.NewProc("viSetBuf")
	procViStatusDesc            = visa32.NewProc("viStatusDesc")
	procViTerminate          = visa32.NewProc("viTerminate")
	procViUninstallHandler   = visa32.NewProc("viUninstallHandler")
	procViUnlock             = visa32.NewProc("viUnlock")
	procViUnmapAddress       = visa32.NewProc("viUnmapAddress")
	procViUnmapTrigger       = visa32.NewProc("viUnmapTrigger")
	procViUsbControlIn       = visa32.NewProc("viUsbControlIn")
	procViUsbControlOut      = visa32.NewProc("viUsbControlOut")
	procViVPrintf            = visa32.NewProc("viVPrintf")
	procViVQueryf            = visa32.NewProc("viVQueryf")
	procViVSPrintf           = visa32.NewProc("viVSPrintf")
	procViVSScanf            = visa32.NewProc("viVSScanf")
	procViVScanf             = visa32.NewProc("viVScanf")
	procViVxiCommandQuery    = visa32.NewProc("viVxiCommandQuery")
	procViWaitOnEvent        = visa32.NewProc("viWaitOnEvent")
	procViWrite              = visa32.NewProc("viWrite")
	procViWriteAsync         = visa32.NewProc("viWriteAsync")
	procViWriteFromFile      = visa32.NewProc("viWriteFromFile")
)

type Visa struct {
	defaultRM uint32
	vi uint32
}

func TrimByteSlice(b []byte) []byte {
	pos := func(c byte) int {
		for p, v := range b {
			if v == c {
				return p
			}
		}
		return len(b)
	}

	return b[0:pos(byte(0))]
}

// only for integer type
func intPtr(n interface{}) uintptr {
	switch value := n.(type) {
	case *int8:
		return uintptr(unsafe.Pointer(value))
	case *int16:
		return uintptr(unsafe.Pointer(value))
	case *int32:
		return uintptr(unsafe.Pointer(value))
	case *int64:
		return uintptr(unsafe.Pointer(value))
	case *uint8:
		return uintptr(unsafe.Pointer(value))
	case *uint16:
		return uintptr(unsafe.Pointer(value))
	case *uint32:
		return uintptr(unsafe.Pointer(value))
	case *uint64:
		return uintptr(unsafe.Pointer(value))
	default:
		return uintptr(unsafe.Pointer(n.(*int)))
	}
}

func strPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringBytePtr(s)))
}

func slicePtr(b []byte) uintptr {
	return uintptr(unsafe.Pointer(&b[0]))
}

func abort(funcname string, _errCode uint) {
	errStr, ok := errorCode[_errCode]
	if ok {
		panic(fmt.Sprintf("%s failed. Error Msg: %s", funcname, errStr))
	} else {
		panic(fmt.Sprintf("%s failed. Error Code: 0x%X", funcname, _errCode))
	}
}

func (v *Visa)ViAssertIntrSignal(mode int16, statusID uint32) int32 {
	return ViAssertIntrSignal(v.vi, mode, statusID)
}

func ViAssertIntrSignal(vi uint32, mode int16, statusID uint32) int32 {
	ret, _, _ := procViAssertIntrSignal.Call(uintptr(vi),
		intPtr(mode),
		uintptr(statusID))
	if int32(ret) < VI_SUCCESS {
		abort("viAssertIntrSignal", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViAssertTrigger(protocol int16) int32 {
	return ViAssertTrigger(v.vi, protocol)
}

func ViAssertTrigger(vi uint32, protocol int16) int32 {
	ret, _, _ := procViAssertTrigger.Call(uintptr(vi), uintptr(protocol))
	if int32(ret) < VI_SUCCESS {
		abort("viAssertTrigger", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViAssertUtilSignal(line int16) int32 {
	return ViAssertUtilSignal(v.vi, line)
}

func ViAssertUtilSignal(vi uint32, line int16) int32 {
	ret, _, _ := procViAssertUtilSignal.Call(uintptr(vi), uintptr(line))
	if int32(ret) < VI_SUCCESS {
		abort("viAssertUtilSignal", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViBufRead(buf []byte, cnt uint32, retCnt *uint32) int32 {
	return ViBufRead(v.vi, buf, cnt, retCnt)
}

func ViBufRead(vi uint32, buf []byte, cnt uint32, retCnt *uint32) int32 {
	ret, _, _ := procViBufRead.Call(uintptr(vi),
		slicePtr(buf),
		uintptr(cnt),
		intPtr(retCnt))
	if int32(ret) < VI_SUCCESS {
		abort("viBufRead", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViBufWrite(buf []byte, cnt uint32, retCnt *uint32) int32 {
	return ViBufWrite(v.vi, buf, cnt, retCnt)
}

func ViBufWrite(vi uint32, buf []byte, cnt uint32, retCnt *uint32) int32 {
	ret, _, _ := procViBufWrite.Call(uintptr(vi),
		slicePtr(buf),
		uintptr(cnt),
		intPtr(retCnt))
	if int32(ret) < VI_SUCCESS {
		abort("viBufWrite", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViClear() int32 {
	return ViClear(v.vi)
}

func ViClear(vi uint32) int32 {
	ret, _, _ := procViClear.Call(uintptr(vi))
	if int32(ret) < VI_SUCCESS {
		abort("viClear", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViClose() int32 {
	return ViClose(v.vi)
}

func ViClose(vi uint32) int32 {
	ret, _, _ := procViClose.Call(uintptr(vi))
	if int32(ret) < VI_SUCCESS {
		abort("viClose", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViDisableEvent(eventType uint32, mechanism uint16) int32 {
	return ViDisableEvent(v.vi, eventType, mechanism)
}

func ViDisableEvent(vi uint32, eventType uint32, mechanism uint16) int32 {
	ret, _, _ := procViDisableEvent.Call(uintptr(vi),
		uintptr(eventType),
		uintptr(mechanism))
	if int32(ret) < VI_SUCCESS {
		abort("viDisableEvent", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViDiscardEvents(eventType uint32, mechanism uint16) int32 {
	return ViDiscardEvents(v.vi, eventType, mechanism)
}

func ViDiscardEvents(vi uint32, eventType uint32, mechanism uint16) int32 {
	ret, _, _ := procViDiscardEvents.Call(uintptr(vi),
		uintptr(eventType),
		uintptr(mechanism))
	if int32(ret) < VI_SUCCESS {
		abort("viDiscardEvents", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViEnableEvent(eventType uint32, mechanism uint16, context uint32) int32 {
	return ViEnableEvent(v.vi, eventType, mechanism, context)
}

func ViEnableEvent(vi uint32, eventType uint32, mechanism uint16, context uint32) int32 {
	ret, _, _ := procViEnableEvent.Call(uintptr(vi),
		uintptr(eventType),
		uintptr(mechanism),
		uintptr(context))
	if int32(ret) < VI_SUCCESS {
		abort("viEnableEvent", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViFindNext(desc string) int32 {
	return ViFindNext(v.vi, desc)
}

func ViFindNext(vi uint32, desc string) int32 {
	ret, _, _ := procViFindNext.Call(uintptr(vi), strPtr(desc))
	if int32(ret) < VI_SUCCESS {
		abort("viFindNext", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViFindRsrc(sesn uint32, expr string, retCnt *uint32, desc []byte) int32 {
	return ViFindRsrc(sesn, expr, &v.vi, retCnt, desc) 
}

func ViFindRsrc(sesn uint32, expr string, vi *uint32, retCnt *uint32, desc []byte) int32 {
	ret, _, _ := procViFindRsrc.Call(uintptr(sesn),
		strPtr(expr),
		intPtr(vi),
		intPtr(retCnt),
		slicePtr(desc))

	if int32(ret) < VI_SUCCESS {
		abort("viFindRsrc", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViFlush(vi uint32, mask uint16) int32 {
	return ViFlush(v.vi, mask)
}

func ViFlush(vi uint32, mask uint16) int32 {
	ret, _, _ := procViFlush.Call(uintptr(vi), uintptr(mask))
	if int32(ret) < VI_SUCCESS {
		abort("viFlush", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViGetAttribute(attrName uint32, attrValue unsafe.Pointer) int32 {
	return ViGetAttribute(v.vi, attrName, attrValue)
}

func ViGetAttribute(vi uint32, attrName uint32, attrValue unsafe.Pointer) int32 {
	ret, _, _ := procViGetAttribute.Call(uintptr(vi), uintptr(attrValue))
	if int32(ret) < VI_SUCCESS {
		abort("viGetAttribute", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViGetDefaultRM() int32 {
	return ViGetDefaultRM(&v.vi)
}

func ViGetDefaultRM(vi *uint32) int32 {
	return ViOpenDefaultRM(vi)
}

func (v *Visa)ViGpibCommand(cmd string, cnt uint32, retCnt *uint32) int32 {
	return ViGpibCommand(v.vi, cmd, cnt, retCnt)
}

func ViGpibCommand(vi uint32, cmd string, cnt uint32, retCnt *uint32) int32 {
	ret, _, _ := procViGpibCommand.Call(uintptr(vi),
		strPtr(cmd),
		uintptr(cnt),
		intPtr(retCnt))
	if int32(ret) < VI_SUCCESS {
		abort("viGpibCommand", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViGpibControlATN(mode uint16) int32 {
	return ViGpibControlATN(v.vi, mode)
}

func ViGpibControlATN(vi uint32, mode uint16) int32 {
	ret, _, _ := procViGpibControlATN.Call(uintptr(vi),
		uintptr(mode))
	if int32(ret) < VI_SUCCESS {
		abort("ViGpibControlATN", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViGpibControlREN(mode uint16) int32 {
	return ViGpibControlREN(v.vi, mode)
}

func ViGpibControlREN(vi uint32, mode uint16) int32 {
	ret, _, _ := procViGpibControlREN.Call(uintptr(vi),
		uintptr(mode))
	if int32(ret) < VI_SUCCESS {
		abort("ViGpibControlREN", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViGpibPassControl(primAddr, secAddr uint16) int32 {
	return ViGpibPassControl(v.vi, primAddr, secAddr)
}

func ViGpibPassControl(vi uint32, primAddr, secAddr uint16) int32 {
	ret, _, _ := procViGpibPassControl.Call(uintptr(vi),
		uintptr(primAddr),
		uintptr(secAddr))
	if int32(ret) < VI_SUCCESS {
		abort("ViGpibPassControl", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViGpibSendIFC() int32 {
	return ViGpibSendIFC(v.vi)
}

func ViGpibSendIFC(vi uint32) int32 {
	ret, _, _ := procViGpibSendIFC.Call(uintptr(vi))
	if int32(ret) < VI_SUCCESS {
		abort("ViGpibSendIFC", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViIn16(space uint16, offset uint, val16 *uint16) int32 {
	return ViIn16(v.vi, space, offset, val16)
}

func ViIn16(vi uint32, space uint16, offset uint, val16 *uint16) int32 {
	ret, _, _ := procViIn16.Call(uintptr(vi), 
		uintptr(space), 
		uintptr(offset), 
		intPtr(val16))
	if int32(ret) < VI_SUCCESS {
		abort("ViIn16", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViIn16Ex(space uint16, offset uint64, val16 *uint16) int32 {
	return ViIn16Ex(v.vi, space, offset, val16)
}

func ViIn16Ex(vi uint32, space uint16, offset uint64, val16 *uint16) int32 {
	ret, _, _ := procViIn16Ex.Call(uintptr(vi), 
		uintptr(space), 
		uintptr(offset), 
		intPtr(val16))
	if int32(ret) < VI_SUCCESS {
		abort("ViIn16Ex", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViIn32(space uint16, offset uint, val32 *uint32) int32 {
	return ViIn32(v.vi, space, offset, val32)
}

func ViIn32(vi uint32, space uint16, offset uint, val32 *uint32) int32 {
	ret, _, _ := procViIn32.Call(uintptr(vi), 
		uintptr(space), 
		uintptr(offset), 
		intPtr(val32))
	if int32(ret) < VI_SUCCESS {
		abort("ViIn32", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViIn32Ex(space uint16, offset uint64, val32 *uint32) int32 {
	return ViIn32Ex(v.vi, space, offset, val32)
}

func ViIn32Ex(vi uint32, space uint16, offset uint64, val32 *uint32) int32 {
	ret, _, _ := procViIn32Ex.Call(uintptr(vi), 
		uintptr(space), 
		uintptr(offset), 
		intPtr(val32))
	if int32(ret) < VI_SUCCESS {
		abort("ViIn32Ex", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViIn64(space uint16, offset uint, val64 *uint64) int32 {
	return ViIn64(v.vi, space, offset, val64)
}

func ViIn64(vi uint32, space uint16, offset uint, val64 *uint64) int32 {
	ret, _, _ := procViIn64.Call(uintptr(vi), 
		uintptr(space), 
		uintptr(offset), 
		intPtr(val64))
	if int32(ret) < VI_SUCCESS {
		abort("ViIn64", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViIn64Ex(space uint16, offset uint64, val64 *uint64) int32 {
	return ViIn64Ex(v.vi, space, offset, val64)
}

func ViIn64Ex(vi uint32, space uint16, offset uint64, val64 *uint64) int32 {
	ret, _, _ := procViIn64Ex.Call(uintptr(vi), 
		uintptr(space), 
		uintptr(offset), 
		intPtr(val64))
	if int32(ret) < VI_SUCCESS {
		abort("ViIn64Ex", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViIn8(space uint16, offset uint, val8 *uint8) int32 {
	return ViIn8(v.vi, space, offset, val8)
}

func ViIn8(vi uint32, space uint16, offset uint, val8 *uint8) int32 {
	ret, _, _ := procViIn8.Call(uintptr(vi), 
		uintptr(space), 
		uintptr(offset), 
		intPtr(val8))
	if int32(ret) < VI_SUCCESS {
		abort("ViIn8", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViIn8Ex(space uint16, offset uint64, val8 *uint8) int32 {
	return ViIn8Ex(v.vi, space, offset, val8)
}

func ViIn8Ex(vi uint32, space uint16, offset uint64, val8 *uint8) int32 {
	ret, _, _ := procViIn8Ex.Call(uintptr(vi), 
		uintptr(space), 
		uintptr(offset), 
		intPtr(val8))
	if int32(ret) < VI_SUCCESS {
		abort("ViIn8Ex", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViInstallHandler(eventType uint32, handler *int32, userHandle unsafe.Pointer) int32 {
	return ViInstallHandler(v.vi, eventType, handler, userHandle)
}

func ViInstallHandler(vi, eventType uint32, handler *int32, userHandle unsafe.Pointer) int32 {
	ret, _, _ := procViInstallHandler.Call(uintptr(vi), 
		uintptr(eventType), 
		intPtr(handler), 
		uintptr(userHandle))
	if int32(ret) < VI_SUCCESS {
		abort("ViInstallHandler", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViLock(lockType, timeout uint32, requestedKey string, accessKey []byte) int32 {
	return ViLock(v.vi, lockType, timeout, requestedKey, accessKey)
}

func ViLock(vi, lockType, timeout uint32, requestedKey string, accessKey []byte) int32 {
	ret, _, _ := procViLock.Call(uintptr(vi), 
		uintptr(lockType), 
		uintptr(timeout), 
		strPtr(requestedKey),
		slicePtr(accessKey))
	if int32(ret) < VI_SUCCESS {
		abort("ViLock", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMapAddress(mapSpace uint16, mapOffset, mapSize uint, 
	access uint16, suggested, address unsafe.Pointer) int32 {
		return ViMapAddress(v.vi, mapSpace, mapOffset, mapSize, access, suggested, address)
	}

func ViMapAddress(vi uint32, mapSpace uint16, mapOffset, 
	mapSize uint, access uint16, suggested, address unsafe.Pointer) int32 {
		ret, _, _ := procViMapAddress.Call(uintptr(vi), 
		uintptr(mapSpace), 
		uintptr(mapOffset), 
		uintptr(mapSize), 
		uintptr(access), 
		uintptr(suggested),
		uintptr(address))
	if int32(ret) < VI_SUCCESS {
		abort("ViMapAddress", uint(ret))
	}
	return int32(ret)
		
}

func (v *Visa)ViMapAddressEx(mapSpace uint16, mapOffset uint64, mapSize uint, 
	access uint16, suggested, address unsafe.Pointer) int32 {
		return ViMapAddressEx(v.vi, mapSpace, mapOffset, mapSize, access, suggested, address)
	}
	
func ViMapAddressEx(vi uint32, mapSpace uint16, mapOffset uint64,
	mapSize uint, access uint16, suggested, address unsafe.Pointer) int32 {
		ret, _, _ := procViMapAddressEx.Call(uintptr(vi), 
		uintptr(mapSpace), 
		uintptr(mapOffset), 
		uintptr(mapSize), 
		uintptr(access), 
		uintptr(suggested),
		uintptr(address))
	if int32(ret) < VI_SUCCESS {
		abort("ViMapAddressEx", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMapTrigger(trigSrc, trigDest int16, mode uint16) int32 {
	return ViMapTrigger(v.vi, trigSrc, trigDest, mode)
}

func ViMapTrigger(vi uint32, trigSrc, trigDest int16, mode uint16) int32 {
	ret, _, _ := procViMapTrigger.Call(uintptr(vi), 
		uintptr(trigSrc), 
		uintptr(trigDest), 
		uintptr(mode))
	if int32(ret) < VI_SUCCESS {
		abort("ViMapTrigger", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMemAlloc(size uint, offset *uint) int32 {
	return ViMemAlloc(v.vi, size, offset)
}

func ViMemAlloc(vi uint32, size uint, offset *uint) int32 {
	ret, _, _ := procViMemAlloc.Call(uintptr(vi), 
		uintptr(size), 
		intPtr(offset))
	if int32(ret) < VI_SUCCESS {
		abort("ViMemAlloc", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMemAllocEx(size uint, offset *uint64) int32 {
	return ViMemAllocEx(v.vi, size, offset)
}

func ViMemAllocEx(vi uint32, size uint, offset *uint64) int32 {
	ret, _, _ := procViMemAllocEx.Call(uintptr(vi), 
		uintptr(size), 
		intPtr(offset))
	if int32(ret) < VI_SUCCESS {
		abort("ViMemAllocEx", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMemFree(offset *uint) int32 {
	return ViMemFree(v.vi, offset)
}

func ViMemFree(vi uint32, offset *uint) int32 {
	ret, _, _ := procViMemFree.Call(uintptr(vi), 
		intPtr(offset))
	if int32(ret) < VI_SUCCESS {
		abort("ViMemFree", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMemFreeEx(offset *uint64) int32 {
	return ViMemFreeEx(v.vi, offset)
}

func ViMemFreeEx(vi uint32, offset *uint64) int32 {
	ret, _, _ := procViMemFreeEx.Call(uintptr(vi), 
		intPtr(offset))
	if int32(ret) < VI_SUCCESS {
		abort("ViMemFreeEx", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMove(srcSpace uint16, srcOffset uint, srcWidth, destSpace uint16,
    destOffset uint64, destWidth uint16, srcLength uint) int32 {
    	return ViMove(v.vi, srcSpace, srcOffset, srcWidth, destSpace,
    destOffset, destWidth, srcLength)
    }

func ViMove(vi uint32, srcSpace uint16, srcOffset uint, srcWidth, destSpace uint16,
    destOffset uint64, destWidth uint16, srcLength uint) int32 {
    	ret, _, _ := procViMove.Call(uintptr(vi), 
		uintptr(srcSpace),
		uintptr(srcOffset),
		uintptr(srcWidth),
		uintptr(destSpace),
		uintptr(destOffset),
		uintptr(destWidth),
		uintptr(srcLength))
    	
	if int32(ret) < VI_SUCCESS {
		abort("ViMove", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMoveAsync(srcSpace uint16, srcOffset uint, srcWidth, destSpace uint16,
    destOffset uint64, destWidth uint16, srcLength uint, jobId *uint32) int32 {
    	return ViMoveAsync(v.vi, srcSpace, srcOffset , srcWidth, destSpace,
    destOffset, destWidth, srcLength, jobId)
    }

func ViMoveAsync(vi uint32, srcSpace uint16, srcOffset uint, srcWidth, destSpace uint16,
    destOffset uint64, destWidth uint16, srcLength uint, jobId *uint32) int32 {
    	ret, _, _ := procViMoveAsync.Call(uintptr(vi), 
		uintptr(srcSpace),
		uintptr(srcOffset),
		uintptr(srcWidth),
		uintptr(destSpace),
		uintptr(destOffset),
		uintptr(destWidth),
		uintptr(srcLength),
		intPtr(jobId))
    	
	if int32(ret) < VI_SUCCESS {
		abort("ViMoveAsync", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMoveAsyncEx(srcSpace uint16, srcOffset uint64, srcWidth, destSpace uint16,
    destOffset uint64, destWidth uint16, srcLength uint, jobId *uint32) int32 {
    	return ViMoveAsyncEx(v.vi, srcSpace, srcOffset , srcWidth, destSpace,
    destOffset, destWidth, srcLength, jobId)
    }
    
func ViMoveAsyncEx(vi uint32, srcSpace uint16, srcOffset uint64, srcWidth, destSpace uint16,
    destOffset uint64, destWidth uint16, srcLength uint, jobId *uint32) int32 {
    	ret, _, _ := procViMoveAsyncEx.Call(uintptr(vi), 
		uintptr(srcSpace),
		uintptr(srcOffset),
		uintptr(srcWidth),
		uintptr(destSpace),
		uintptr(destOffset),
		uintptr(destWidth),
		uintptr(srcLength),
		intPtr(jobId))
    	
	if int32(ret) < VI_SUCCESS {
		abort("ViMoveAsyncEx", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMoveEx(srcSpace uint16, srcOffset uint64, srcWidth, destSpace uint16,
    destOffset uint64, destWidth uint16, srcLength uint) int32 {
    	return ViMoveEx(v.vi, srcSpace, srcOffset, srcWidth, destSpace,
    destOffset, destWidth, srcLength)
    }
    
func ViMoveEx(vi uint32, srcSpace uint16, srcOffset uint64, srcWidth, destSpace uint16,
    destOffset uint64, destWidth uint16, srcLength uint) int32 {
    	ret, _, _ := procViMoveEx.Call(uintptr(vi), 
		uintptr(srcSpace),
		uintptr(srcOffset),
		uintptr(srcWidth),
		uintptr(destSpace),
		uintptr(destOffset),
		uintptr(destWidth),
		uintptr(srcLength))
    	
	if int32(ret) < VI_SUCCESS {
		abort("ViMoveEx", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMoveIn16(space uint16, offset *uint, length uint32, buf16 *uint16) int32 {
	return ViMoveIn16(v.vi, space, offset, length, buf16)
}

func ViMoveIn16(vi uint32, space uint16, offset *uint, length uint32, buf16 *uint16) int32 {
	ret, _, _ := procViMoveIn16.Call(uintptr(vi), 
		uintptr(space),
		intPtr(offset),
		uintptr(length),
		intPtr(buf16))
    	
	if int32(ret) < VI_SUCCESS {
		abort("ViMoveIn16", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMoveIn16Ex(space uint16, offset *uint64, length uint32, buf16 *uint16) int32 {
	return ViMoveIn16Ex(v.vi, space, offset, length, buf16)
}

func ViMoveIn16Ex(vi uint32, space uint16, offset *uint64, length uint32, buf16 *uint16) int32 {
	ret, _, _ := procViMoveIn16Ex.Call(uintptr(vi), 
		uintptr(space),
		intPtr(offset),
		uintptr(length),
		intPtr(buf16))
    	
	if int32(ret) < VI_SUCCESS {
		abort("ViMoveIn16Ex", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMoveIn32(space uint16, offset *uint, length uint32, buf32 *uint32) int32 {
	return ViMoveIn32(v.vi, space, offset, length, buf32)
}

func ViMoveIn32(vi uint32, space uint16, offset *uint, length uint32, buf32 *uint32) int32 {
	ret, _, _ := procViMoveIn32.Call(uintptr(vi), 
		uintptr(space),
		intPtr(offset),
		uintptr(length),
		intPtr(buf32))
    	
	if int32(ret) < VI_SUCCESS {
		abort("ViMoveIn32", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMoveIn32Ex(space uint16, offset *uint64, length uint32, buf32 *uint32) int32 {
	return ViMoveIn32Ex(v.vi, space, offset, length, buf32)
}

func ViMoveIn32Ex(vi uint32, space uint16, offset *uint64, length uint32, buf32 *uint32) int32 {
	ret, _, _ := procViMoveIn32Ex.Call(uintptr(vi), 
		uintptr(space),
		intPtr(offset),
		uintptr(length),
		intPtr(buf32))
    	
	if int32(ret) < VI_SUCCESS {
		abort("ViMoveIn32Ex", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMoveIn64(space uint16, offset *uint, length uint32, buf64 *uint64) int32 {
	return ViMoveIn64(v.vi, space, offset, length, buf64)
}

func ViMoveIn64(vi uint32, space uint16, offset *uint, length uint32, buf64 *uint64) int32 {
	ret, _, _ := procViMoveIn64.Call(uintptr(vi), 
		uintptr(space),
		intPtr(offset),
		uintptr(length),
		intPtr(buf64))
    	
	if int32(ret) < VI_SUCCESS {
		abort("ViMoveIn64", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMoveIn64Ex(space uint16, offset *uint64, length uint32, buf64 *uint64) int32 {
	return ViMoveIn64Ex(v.vi, space, offset, length, buf64)
}

func ViMoveIn64Ex(vi uint32, space uint16, offset *uint64, length uint32, buf64 *uint64) int32 {
	ret, _, _ := procViMoveIn64Ex.Call(uintptr(vi), 
		uintptr(space),
		intPtr(offset),
		uintptr(length),
		intPtr(buf64))
    	
	if int32(ret) < VI_SUCCESS {
		abort("ViMoveIn64Ex", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMoveIn8(space uint16, offset *uint, length uint32, buf8 *uint8) int32 {
	return ViMoveIn8(v.vi, space, offset, length, buf8)
}

func ViMoveIn8(vi uint32, space uint16, offset *uint, length uint32, buf8 *uint8) int32 {
	ret, _, _ := procViMoveIn8.Call(uintptr(vi), 
		uintptr(space),
		intPtr(offset),
		uintptr(length),
		intPtr(buf8))
    	
	if int32(ret) < VI_SUCCESS {
		abort("ViMoveIn8", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMoveIn8Ex(space uint16, offset *uint64, length uint32, buf8 *uint8) int32 {
	return ViMoveIn8Ex(v.vi, space, offset, length, buf8)
}

func ViMoveIn8Ex(vi uint32, space uint16, offset *uint64, length uint32, buf8 *uint8) int32 {
	ret, _, _ := procViMoveIn8Ex.Call(uintptr(vi), 
		uintptr(space),
		intPtr(offset),
		uintptr(length),
		intPtr(buf8))
    	
	if int32(ret) < VI_SUCCESS {
		abort("ViMoveIn8Ex", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMoveOut16(space uint16, offset *uint, length uint32, buf16 *uint16) int32 {
	return ViMoveOut16(v.vi, space, offset, length, buf16)
}

func ViMoveOut16(vi uint32, space uint16, offset *uint, length uint32, buf16 *uint16) int32 {
	ret, _, _ := procViMoveOut16.Call(uintptr(vi), 
		uintptr(space),
		intPtr(offset),
		uintptr(length),
		intPtr(buf16))
    	
	if int32(ret) < VI_SUCCESS {
		abort("ViMoveOut16", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMoveOut16Ex(space uint16, offset *uint64, length uint32, buf16 *uint16) int32 {
	return ViMoveOut16Ex(v.vi, space, offset, length, buf16)
}

func ViMoveOut16Ex(vi uint32, space uint16, offset *uint64, length uint32, buf16 *uint16) int32 {
	ret, _, _ := procViMoveOut16Ex.Call(uintptr(vi), 
		uintptr(space),
		intPtr(offset),
		uintptr(length),
		intPtr(buf16))
    	
	if int32(ret) < VI_SUCCESS {
		abort("ViMoveOut16Ex", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMoveOut32(space uint16, offset *uint, length uint32, buf32 *uint32) int32 {
	return ViMoveOut32(v.vi, space, offset, length, buf32)
}

func ViMoveOut32(vi uint32, space uint16, offset *uint, length uint32, buf32 *uint32) int32 {
	ret, _, _ := procViMoveOut32.Call(uintptr(vi), 
		uintptr(space),
		intPtr(offset),
		uintptr(length),
		intPtr(buf32))
    	
	if int32(ret) < VI_SUCCESS {
		abort("ViMoveOut32", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMoveOut32Ex(space uint16, offset *uint64, length uint32, buf32 *uint32) int32 {
	return ViMoveOut32Ex(v.vi, space, offset, length, buf32)
}

func ViMoveOut32Ex(vi uint32, space uint16, offset *uint64, length uint32, buf32 *uint32) int32 {
	ret, _, _ := procViMoveOut32Ex.Call(uintptr(vi), 
		uintptr(space),
		intPtr(offset),
		uintptr(length),
		intPtr(buf32))
    	
	if int32(ret) < VI_SUCCESS {
		abort("ViMoveOut32Ex", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMoveOut64(space uint16, offset *uint, length uint32, buf64 *uint64) int32 {
	return ViMoveOut64(v.vi, space, offset, length, buf64)
}

func ViMoveOut64(vi uint32, space uint16, offset *uint, length uint32, buf64 *uint64) int32 {
	ret, _, _ := procViMoveOut64.Call(uintptr(vi), 
		uintptr(space),
		intPtr(offset),
		uintptr(length),
		intPtr(buf64))
    	
	if int32(ret) < VI_SUCCESS {
		abort("ViMoveOut64", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMoveOut64Ex(space uint16, offset *uint64, length uint32, buf64 *uint64) int32 {
	return ViMoveOut64Ex(v.vi, space, offset, length, buf64)
}

func ViMoveOut64Ex(vi uint32, space uint16, offset *uint64, length uint32, buf64 *uint64) int32 {
	ret, _, _ := procViMoveOut64Ex.Call(uintptr(vi), 
		uintptr(space),
		intPtr(offset),
		uintptr(length),
		intPtr(buf64))
    	
	if int32(ret) < VI_SUCCESS {
		abort("ViMoveOut64Ex", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMoveOut8(space uint16, offset *uint, length uint32, buf8 *uint8) int32 {
	return ViMoveOut8(v.vi, space, offset, length, buf8)
}

func ViMoveOut8(vi uint32, space uint16, offset *uint, length uint32, buf8 *uint8) int32 {
	ret, _, _ := procViMoveOut8.Call(uintptr(vi), 
		uintptr(space),
		intPtr(offset),
		uintptr(length),
		intPtr(buf8))
    	
	if int32(ret) < VI_SUCCESS {
		abort("ViMoveOut8", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViMoveOut8Ex(space uint16, offset *uint64, length uint32, buf8 *uint8) int32 {
	return ViMoveOut8Ex(v.vi, space, offset, length, buf8)
}

func ViMoveOut8Ex(vi uint32, space uint16, offset *uint64, length uint32, buf8 *uint8) int32 {
	ret, _, _ := procViMoveOut8Ex.Call(uintptr(vi), 
		uintptr(space),
		intPtr(offset),
		uintptr(length),
		intPtr(buf8))
    	
	if int32(ret) < VI_SUCCESS {
		abort("ViMoveOut8Ex", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViOpen(name string, mode uint32, timeout uint32) int32 {
	return ViOpen(v.defaultRM, name, mode, timeout, &v.vi)
}

func ViOpen(sesn uint32, name string, mode uint32, timeout uint32, vi *uint32) int32 {
	ret, _, _ := procViOpen.Call(uintptr(sesn),
		strPtr(name),
		uintptr(mode),
		uintptr(timeout),
		intPtr(vi))
	if int32(ret) < VI_SUCCESS {
		abort("viOpen", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViOpenDefaultRM() int32 {
	return ViOpenDefaultRM(&v.defaultRM)
}

func ViOpenDefaultRM(vi *uint32) int32 {
	ret, _, _ := procViOpenDefaultRM.Call(intPtr(vi))
	if int32(ret) < VI_SUCCESS {
		abort("viOpenDefaultRM", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViOut16(space uint16, offset uint, val16 uint16) int32 {
	return ViOut16(v.vi, space, offset, val16)
}

func ViOut16(vi uint32, space uint16, offset uint, val16 uint16) int32 {
	ret, _, _ := procViOut16.Call(uintptr(vi), 
		uintptr(space), 
		uintptr(offset), 
		uintptr(val16))
	if int32(ret) < VI_SUCCESS {
		abort("ViOut16", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViOut16Ex(space uint16, offset uint64, val16 uint16) int32 {
	return ViOut16Ex(v.vi, space, offset, val16)
}

func ViOut16Ex(vi uint32, space uint16, offset uint64, val16 uint16) int32 {
	ret, _, _ := procViOut16Ex.Call(uintptr(vi), 
		uintptr(space), 
		uintptr(offset), 
		uintptr(val16))
	if int32(ret) < VI_SUCCESS {
		abort("ViOut16Ex", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViOut32(space uint16, offset uint, val32 uint32) int32 {
	return ViOut32(v.vi, space, offset, val32)
}

func ViOut32(vi uint32, space uint16, offset uint, val32 uint32) int32 {
	ret, _, _ := procViOut32.Call(uintptr(vi), 
		uintptr(space), 
		uintptr(offset), 
		uintptr(val32))
	if int32(ret) < VI_SUCCESS {
		abort("ViOut32", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViOut32Ex(space uint16, offset uint64, val32 uint32) int32 {
	return ViOut32Ex(v.vi, space, offset, val32)
}

func ViOut32Ex(vi uint32, space uint16, offset uint64, val32 uint32) int32 {
	ret, _, _ := procViOut32Ex.Call(uintptr(vi), 
		uintptr(space), 
		uintptr(offset), 
		uintptr(val32))
	if int32(ret) < VI_SUCCESS {
		abort("ViOut32Ex", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViOut64(space uint16, offset uint, val64 uint64) int32 {
	return ViOut64(v.vi, space, offset, val64)
}

func ViOut64(vi uint32, space uint16, offset uint, val64 uint64) int32 {
	ret, _, _ := procViOut64.Call(uintptr(vi), 
		uintptr(space), 
		uintptr(offset), 
		uintptr(val64))
	if int32(ret) < VI_SUCCESS {
		abort("ViOut64", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViOut64Ex(space uint16, offset uint64, val64 uint64) int32 {
	return ViOut64Ex(v.vi, space, offset, val64)
}

func ViOut64Ex(vi uint32, space uint16, offset uint64, val64 uint64) int32 {
	ret, _, _ := procViOut64Ex.Call(uintptr(vi), 
		uintptr(space), 
		uintptr(offset), 
		uintptr(val64))
	if int32(ret) < VI_SUCCESS {
		abort("ViOut64Ex", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViOut8(space uint16, offset uint, val8 uint8) int32 {
	return ViOut8(v.vi, space, offset, val8)
}

func ViOut8(vi uint32, space uint16, offset uint, val8 uint8) int32 {
	ret, _, _ := procViOut8.Call(uintptr(vi), 
		uintptr(space), 
		uintptr(offset), 
		uintptr(val8))
	if int32(ret) < VI_SUCCESS {
		abort("ViOut8", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViOut8Ex(space uint16, offset uint64, val8 uint8) int32 {
	return ViOut8Ex(v.vi, space, offset, val8)
}

func ViOut8Ex(vi uint32, space uint16, offset uint64, val8 uint8) int32 {
	ret, _, _ := procViOut8Ex.Call(uintptr(vi), 
		uintptr(space), 
		uintptr(offset), 
		uintptr(val8))
	if int32(ret) < VI_SUCCESS {
		abort("ViOut8Ex", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViParseRsrc(rsrcName string, intfType *uint16, intfNum *uint16) int32 {
	return ViParseRsrc(v.defaultRM, rsrcName, intfType, intfNum)
}

func ViParseRsrc(rmSesn uint32, rsrcName string, intfType *uint16, intfNum *uint16) int32 {
	ret, _, _ := procViParseRsrc.Call(uintptr(rmSesn),
		strPtr(rsrcName),
		intPtr(intfType),
		intPtr(intfNum))
	if int32(ret) < VI_SUCCESS {
		abort("viParseRsrc", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViParseRsrcEx(rsrcName string, intfType *uint16, intfNum *uint16, 
	rsrcClass, expandedUnaliasedName, aliasIfExists []byte) int32 {
	return ViParseRsrcEx(v.defaultRM, rsrcName, intfType, intfNum,
		rsrcClass, expandedUnaliasedName, aliasIfExists)
}

func ViParseRsrcEx(rmSesn uint32, rsrcName string, intfType *uint16,
	intfNum *uint16, rsrcClass, expandedUnaliasedName, aliasIfExists []byte) int32 {
	ret, _, _ := procViParseRsrcEx.Call(uintptr(rmSesn),
		strPtr(rsrcName),
		intPtr(intfType),
		intPtr(intfNum),
		slicePtr(rsrcClass),
		slicePtr(expandedUnaliasedName),
		slicePtr(aliasIfExists))
	if int32(ret) < VI_SUCCESS {
		abort("viParseRsrcEx", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViPeek16(address unsafe.Pointer, val16 *uint16) {
	ViPeek16(v.vi, address, val16)
}

func ViPeek16(vi uint32, address unsafe.Pointer, val16 *uint16) {
	procViPeek16.Call(uintptr(vi),
		uintptr(address),
		intPtr(val16))
}

func (v *Visa)ViPeek32(address unsafe.Pointer, val32 *uint32) {
	ViPeek32(v.vi, address, val32)
}

func ViPeek32(vi uint32, address unsafe.Pointer, val32 *uint32) {
	procViPeek32.Call(uintptr(vi),
		uintptr(address),
		intPtr(val32))
}

func (v *Visa)ViPeek64(address unsafe.Pointer, val64 *uint64) {
	ViPeek64(v.vi, address, val64)
}

func ViPeek64(vi uint32, address unsafe.Pointer, val64 *uint64) {
	procViPeek64.Call(uintptr(vi),
		uintptr(address),
		intPtr(val64))
}

func (v *Visa)ViPeek8(address unsafe.Pointer, val8 *uint8) {
	ViPeek8(v.vi, address, val8)
}

func ViPeek8(vi uint32, address unsafe.Pointer, val8 *uint8) {
	_, _, _ = procViPeek8.Call(uintptr(vi),
		uintptr(address),
		intPtr(val8))
}

func (v *Visa)ViPoke16(address unsafe.Pointer, val16 uint16) {
	ViPoke16(v.vi, address, val16)
}

func ViPoke16(vi uint32, address unsafe.Pointer, val16 uint16) {
	_, _, _ = procViPoke16.Call(uintptr(vi),
		uintptr(address),
		uintptr(val16))
}

func (v *Visa)ViPoke32(address unsafe.Pointer, val32 uint32) {
	ViPoke32(v.vi, address, val32)
}

func ViPoke32(vi uint32, address unsafe.Pointer, val32 uint32) {
	_, _, _ = procViPoke32.Call(uintptr(vi),
		uintptr(address),
		uintptr(val32))
}

func (v *Visa)ViPoke64(address unsafe.Pointer, val64 uint64) {
	ViPoke64(v.vi, address, val64)
}

func ViPoke64(vi uint32, address unsafe.Pointer, val64 uint64) {
	_, _, _ = procViPoke64.Call(uintptr(vi),
		uintptr(address),
		intPtr(val64))
}

func (v *Visa)ViPoke8(address unsafe.Pointer, val8 uint8) {
	ViPoke8(v.vi, address, val8)
}

func ViPoke8(vi uint32, address unsafe.Pointer, val8 uint8) {
	_, _, _ = procViPoke8.Call(uintptr(vi),
		uintptr(address),
		uintptr(val8))
}

func (v *Visa)ViPrintf(writeFmt string, args ...interface{}) int32 {
	return ViPrintf(v.vi, writeFmt, args...)
}

func ViPrintf(vi uint32, writeFmt string, args ...interface{}) int32 {
	str := fmt.Sprintf(writeFmt, args...)
	ret, _, _ := procViPrintf.Call(uintptr(vi),
		strPtr("%s"),
		strPtr(str))
	if int32(ret) < VI_SUCCESS {
		abort("viPrintf", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViPxiReserveTriggers(cnt int16, trigBuses, trigLines, failureIndex  *int16) int32 {
	return ViPxiReserveTriggers(v.vi, cnt, trigBuses, trigLines, failureIndex)
}

func ViPxiReserveTriggers(vi uint32, cnt int16, trigBuses, trigLines, failureIndex  *int16) int32 {
	ret, _, _ := procViPxiReserveTriggers.Call(uintptr(vi),
		uintptr(cnt),
		intPtr(trigBuses),
		intPtr(trigLines),
		intPtr(failureIndex))
	if int32(ret) < VI_SUCCESS {
		abort("ViPxiReserveTriggers", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViQueryf(cmd string, go_readFmt string, args ...interface{}) int32 {
	return ViQueryf(v.vi, cmd, go_readFmt, args...)
}

func ViQueryf(vi uint32, cmd string, go_readFmt string, args ...interface{}) int32 {
	buf := make([]byte, 1024)
	ret, _, _ := procViQueryf.Call(uintptr(vi),
		strPtr(cmd),
		strPtr("%t"),
		slicePtr(buf))		
	if int32(ret) < VI_SUCCESS {
		abort("ViQueryf", uint(ret))
	}
	fmt.Println("test", string(TrimByteSlice(buf)))
	fmt.Printf("%T, %v\n", string(TrimByteSlice(buf)), string(TrimByteSlice(buf)))
	num, err := fmt.Sscanf(string(TrimByteSlice(buf)), go_readFmt, args...)
	fmt.Println(num)
	if err != nil {
		panic("ViQueryf failed. Error Msg: " + err.Error())
	}
	return int32(ret)
}

func (v *Visa)ViRead(buf []byte, cnt uint32, retCnt *uint32) int32 {
	return ViRead(v.vi, buf, cnt, retCnt)
}

func ViRead(vi uint32, buf []byte, cnt uint32, retCnt *uint32) int32 {
	ret, _, _ := procViRead.Call(uintptr(vi),
		slicePtr(buf),
		uintptr(cnt),
		intPtr(retCnt))
	if int32(ret) < VI_SUCCESS {
		abort("viRead", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViReadAsync(buf []byte, cnt uint32, jobId *uint32) int32 {
	return ViReadAsync(v.vi, buf, cnt, jobId)
}

func ViReadAsync(vi uint32, buf []byte, cnt uint32, jobId *uint32) int32 {
	ret, _, _ := procViReadAsync.Call(uintptr(vi),
		slicePtr(buf),
		uintptr(cnt),
		intPtr(jobId))
	if int32(ret) < VI_SUCCESS {
		abort("ViReadAsync", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViReadSTB(status *uint16) int32 {
	return ViReadSTB(v.vi, status)
}

func ViReadSTB(vi uint32, status *uint16) int32 {
	ret, _, _ := procViReadSTB.Call(uintptr(vi),
		intPtr(status))
	if int32(ret) < VI_SUCCESS {
		abort("ViReadSTB", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViReadToFile(filename string, cnt uint32, retCnt *uint32) int32 {
	return ViReadToFile(v.vi, filename, cnt, retCnt)
}

func ViReadToFile(vi uint32, filename string, cnt uint32, retCnt *uint32) int32 {
	ret, _, _ := procViReadToFile.Call(uintptr(vi),
		strPtr(filename),
		uintptr(cnt),
		intPtr(retCnt))
	if int32(ret) < VI_SUCCESS {
		abort("ViReadToFile", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViSPrintf(buf *string, writeFmt string, args ...interface{}) int32 {
	return ViSPrintf(v.vi, buf, writeFmt, args...)
}

func ViSPrintf(vi uint32, buf *string, writeFmt string, args ...interface{}) int32 {
//	str := fmt.Sprintf(string(writeFmt), args)
//	ret, _, _ := procViSPrintf.Call(uintptr(vi),
//		slicePtr(buf),
//		strPtr("%s"),
//		strPtr(str))
//	if int32(ret) < VI_SUCCESS {
//		abort("ViSPrintf", uint(ret))
//	}
//	return int32(ret)
	
	// Don't use the vi
	_ = vi
	*buf = fmt.Sprintf(writeFmt, args...)
	return VI_SUCCESS
}

func (v *Visa)ViSScanf(buf *string, readFmt string, args ...interface{}) int32 {
	return ViSScanf(v.vi, buf, readFmt, args...)
}

func ViSScanf(vi uint32, buf *string, readFmt string, args ...interface{}) int32 {
	// Don't use the vi
	_ = vi
	ret, err := fmt.Sscanf(*buf, readFmt, args...)
	if err != nil {
		panic("ViSScanf Failed. Error Msg: " + err.Error())
	}
	return int32(ret)
}

func (v *Visa)ViScanf(readFmt string, args ...interface{}) int32 {
	return ViScanf(v.vi, readFmt, args...)
}

func ViScanf(vi uint32, readFmt string, args ...interface{}) int32 {
	buf := make([]byte, 1024)
	ret, _, _ := procViScanf.Call(uintptr(vi),
		strPtr("%t"),
		slicePtr(buf))
	if int32(ret) < VI_SUCCESS {
		abort("ViScanf", uint(ret))
	}
	
	_, err := fmt.Sscanf(string(TrimByteSlice(buf)), readFmt, args...)
	if err != nil {
		panic("ViScanf failed. Error Msg: " + err.Error())
	}
	
	return int32(ret)
}

func (v *Visa)ViSetAttribute(attrName uint32, attrValue uint32) int32 {
	return ViSetAttribute(v.vi, attrName, attrValue) 
}

func ViSetAttribute(vi uint32, attrName uint32, attrValue uint32) int32 {
	ret, _, _ := procViSetAttribute.Call(uintptr(vi),
		uintptr(attrName),
		uintptr(attrValue))
	if int32(ret) < VI_SUCCESS {
		abort("viSetAttribute", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViSetBuf(mask uint16, size uint32) int32 {
	return ViSetBuf(v.vi, mask, size)
}

func ViSetBuf(vi uint32, mask uint16, size uint32) int32 {
	ret, _, _ := procViSetBuf.Call(uintptr(vi),
		uintptr(mask),
		uintptr(size))
	if int32(ret) < VI_SUCCESS {
		abort("viSetBuf", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViStatusDesc(status int32, desc []byte) int32 {
	return ViStatusDesc(v.vi, status, desc)
}

func ViStatusDesc(vi uint32, status int32, desc []byte) int32 {
	ret, _, _ := procViStatusDesc.Call(uintptr(vi),
		uintptr(status),
		slicePtr(desc))
	if int32(ret) < VI_SUCCESS {
		abort("ViStatusDesc", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViTerminate(degree uint16, jobId uint32) int32 {
	return ViTerminate(v.vi, degree, jobId)
}

func ViTerminate(vi uint32, degree uint16, jobId uint32) int32 {
	ret, _, _ := procViTerminate.Call(uintptr(vi),
		uintptr(degree),
		uintptr(jobId))
	if int32(ret) < VI_SUCCESS {
		abort("ViTerminate", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViUninstallHandler(eventType uint32, handler *int32, userHandle unsafe.Pointer) int32 {
	return ViUninstallHandler(v.vi, eventType, handler, userHandle)
}

func ViUninstallHandler(vi, eventType uint32, handler *int32, userHandle unsafe.Pointer) int32 {
	ret, _, _ := procViUninstallHandler.Call(uintptr(vi),
		uintptr(eventType),
		intPtr(handler),
		uintptr(userHandle))
	if int32(ret) < VI_SUCCESS {
		abort("ViUninstallHandler", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViUnlock() int32 {
	return ViUnlock(v.vi)
}

func ViUnlock(vi uint32) int32 {
	ret, _, _ := procViUnlock.Call(uintptr(vi))
	if int32(ret) < VI_SUCCESS {
		abort("ViUnlock", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViUnmapAddress() int32 {
	return ViUnmapAddress(v.vi)
}

func ViUnmapAddress(vi uint32) int32 {
	ret, _, _ := procViUnmapAddress.Call(uintptr(vi))
	if int32(ret) < VI_SUCCESS {
		abort("ViUnmapAddress", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViUnmapTrigger(trigSrc, trigDest int16) int32 {
	return ViUnmapTrigger(v.vi, trigSrc, trigDest)
}

func ViUnmapTrigger(vi uint32, trigSrc, trigDest int16) int32 {
	ret, _, _ := procViUnmapTrigger.Call(uintptr(vi), uintptr(trigSrc), uintptr(trigDest))
	if int32(ret) < VI_SUCCESS {
		abort("ViUnmapTrigger", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViUsbControlIn(bmRequestType, bRequest int16, wValue, wIndex, wLength uint16, 
	buf []byte, retCnt *uint16) int32 {
		return ViUsbControlIn(v.vi, bmRequestType, bRequest, wValue, wIndex, wLength, buf, retCnt)
	}

func ViUsbControlIn(vi uint32, bmRequestType, bRequest int16, wValue, wIndex, wLength uint16, 
	buf []byte, retCnt *uint16) int32 {
		ret, _, _ := procViUsbControlIn.Call(uintptr(vi), uintptr(bmRequestType), uintptr(bRequest),
			uintptr(wValue), uintptr(wIndex), uintptr(wLength), slicePtr(buf), intPtr(retCnt))
	if int32(ret) < VI_SUCCESS {
		abort("ViUsbControlIn", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViUsbControlOut(bmRequestType, bRequest int16, wValue, wIndex, wLength uint16, buf string) int32 {
	return ViUsbControlOut(v.vi, bmRequestType, bRequest, wValue, wIndex, wLength, buf)
}

func ViUsbControlOut(vi uint32, bmRequestType, bRequest int16, wValue, wIndex, wLength uint16, buf string) int32 {
	ret, _, _ := procViUsbControlOut.Call(uintptr(vi), uintptr(bmRequestType), uintptr(bRequest),
		uintptr(wValue), uintptr(wIndex), uintptr(wLength), strPtr(buf))
	if int32(ret) < VI_SUCCESS {
		abort("ViUsbControlOut", uint(ret))
	}
	return int32(ret)
}

//func ViVPrintf() int32 {
//}
//
//func ViVQueryf() int32 {
//}
//
//func ViVSPrintf() int32 {
//}
//
//func ViVSScanf() int32 {
//}
//
// Need to verify
func (v *Visa)ViVScanf(readFmt string, params []interface{}) int32 {
	return ViVScanf(v.vi, readFmt, params)
}

func ViVScanf(vi uint32, readFmt string, params []interface{}) int32 {
	ret, _, _ := procViVScanf.Call(uintptr(vi),
		strPtr(readFmt),
		uintptr(unsafe.Pointer(&params[0])))
	if int32(ret) < VI_SUCCESS {
		abort("ViVScanf", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViVxiCommandQuery(mode uint16, cmd uint32, response *uint32) int32 {
	return ViVxiCommandQuery(v.vi, mode, cmd, response)
}

func ViVxiCommandQuery(vi uint32, mode uint16, cmd uint32, response *uint32) int32 {
	ret, _, _ := procViVxiCommandQuery.Call(uintptr(vi),
		uintptr(mode),
		uintptr(cmd),
		intPtr(response))
	if int32(ret) < VI_SUCCESS {
		abort("ViVxiCommandQuery", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViWaitOnEvent(inEventType, timeout uint32, outEventType, outContext *uint32) int32 {
	return ViWaitOnEvent(v.vi, inEventType, timeout, outEventType, outContext)
}

func ViWaitOnEvent(vi, inEventType, timeout uint32, outEventType, outContext *uint32) int32 {
	ret, _, _ := procViWaitOnEvent.Call(uintptr(vi),
		uintptr(inEventType),
		uintptr(timeout),
		intPtr(outEventType),
		intPtr(outContext))
	if int32(ret) < VI_SUCCESS {
		abort("ViWaitOnEvent", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViWrite(buf string, cnt uint32, retCnt *uint32) int32 {
	return ViWrite(v.vi, buf, cnt, retCnt)
}

func ViWrite(vi uint32, buf string, cnt uint32, retCnt *uint32) int32 {
	ret, _, _ := procViWrite.Call(uintptr(vi),
		strPtr(buf),
		uintptr(cnt),
		intPtr(retCnt))
	if int32(ret) < VI_SUCCESS {
		abort("viWrite", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViWriteAsync(buf string, cnt uint32, jobId *uint32) int32 {
	return ViWriteAsync(v.vi, buf, cnt, jobId)
}

func ViWriteAsync(vi uint32, buf string, cnt uint32, jobId *uint32) int32 {
	ret, _, _ := procViWriteAsync.Call(uintptr(vi),
		strPtr(buf),
		uintptr(cnt),
		intPtr(jobId))
	if int32(ret) < VI_SUCCESS {
		abort("ViWriteAsync", uint(ret))
	}
	return int32(ret)
}

func (v *Visa)ViWriteFromFile(filename string, cnt uint32, retCnt *uint32) int32 {
	return ViWriteFromFile(v.vi, filename, cnt, retCnt)
}

func ViWriteFromFile(vi uint32, filename string, cnt uint32, retCnt *uint32) int32 {
	ret, _, _ := procViWriteFromFile.Call(uintptr(vi),
		strPtr(filename),
		uintptr(cnt),
		intPtr(retCnt))
	if int32(ret) < VI_SUCCESS {
		abort("ViWriteFromFile", uint(ret))
	}
	return int32(ret)
}
