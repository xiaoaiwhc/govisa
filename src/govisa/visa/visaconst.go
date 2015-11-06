package visa

import (
	"runtime"
	"strings"
)

const (
	VI_SUCCESS = 0
	VI_NULL = 0
	VI_TRUE = true
	VI_FALSE = false
	)

/*- Attributes (platform independent size) ----------------------------------*/
const (
	VI_ATTR_RSRC_CLASS                  uint32 = 0xBFFF0001
	VI_ATTR_RSRC_NAME                   uint32 = 0xBFFF0002
	VI_ATTR_RSRC_IMPL_VERSION           uint32 = 0x3FFF0003
	VI_ATTR_RSRC_LOCK_STATE             uint32 = 0x3FFF0004
	VI_ATTR_MAX_QUEUE_LENGTH            uint32 = 0x3FFF0005
	VI_ATTR_USER_DATA_32                uint32 = 0x3FFF0007
	VI_ATTR_FDC_CHNL                    uint32 = 0x3FFF000D
	VI_ATTR_FDC_MODE                    uint32 = 0x3FFF000F
	VI_ATTR_FDC_GEN_SIGNAL_EN           uint32 = 0x3FFF0011
	VI_ATTR_FDC_USE_PAIR                uint32 = 0x3FFF0013
	VI_ATTR_SEND_END_EN                 uint32 = 0x3FFF0016
	VI_ATTR_TERMCHAR                    uint32 = 0x3FFF0018
	VI_ATTR_TMO_VALUE                   uint32 = 0x3FFF001A
	VI_ATTR_GPIB_READDR_EN              uint32 = 0x3FFF001B
	VI_ATTR_IO_PROT                     uint32 = 0x3FFF001C
	VI_ATTR_DMA_ALLOW_EN                uint32 = 0x3FFF001E
	VI_ATTR_ASRL_BAUD                   uint32 = 0x3FFF0021
	VI_ATTR_ASRL_DATA_BITS              uint32 = 0x3FFF0022
	VI_ATTR_ASRL_PARITY                 uint32 = 0x3FFF0023
	VI_ATTR_ASRL_STOP_BITS              uint32 = 0x3FFF0024
	VI_ATTR_ASRL_FLOW_CNTRL             uint32 = 0x3FFF0025
	VI_ATTR_RD_BUF_OPER_MODE            uint32 = 0x3FFF002A
	VI_ATTR_RD_BUF_SIZE                 uint32 = 0x3FFF002B
	VI_ATTR_WR_BUF_OPER_MODE            uint32 = 0x3FFF002D
	VI_ATTR_WR_BUF_SIZE                 uint32 = 0x3FFF002E
	VI_ATTR_SUPPRESS_END_EN             uint32 = 0x3FFF0036
	VI_ATTR_TERMCHAR_EN                 uint32 = 0x3FFF0038
	VI_ATTR_DEST_ACCESS_PRIV            uint32 = 0x3FFF0039
	VI_ATTR_DEST_BYTE_ORDER             uint32 = 0x3FFF003A
	VI_ATTR_SRC_ACCESS_PRIV             uint32 = 0x3FFF003C
	VI_ATTR_SRC_BYTE_ORDER              uint32 = 0x3FFF003D
	VI_ATTR_SRC_INCREMENT               uint32 = 0x3FFF0040
	VI_ATTR_DEST_INCREMENT              uint32 = 0x3FFF0041
	VI_ATTR_WIN_ACCESS_PRIV             uint32 = 0x3FFF0045
	VI_ATTR_WIN_BYTE_ORDER              uint32 = 0x3FFF0047
	VI_ATTR_GPIB_ATN_STATE              uint32 = 0x3FFF0057
	VI_ATTR_GPIB_ADDR_STATE             uint32 = 0x3FFF005C
	VI_ATTR_GPIB_CIC_STATE              uint32 = 0x3FFF005E
	VI_ATTR_GPIB_NDAC_STATE             uint32 = 0x3FFF0062
	VI_ATTR_GPIB_SRQ_STATE              uint32 = 0x3FFF0067
	VI_ATTR_GPIB_SYS_CNTRL_STATE        uint32 = 0x3FFF0068
	VI_ATTR_GPIB_HS488_CBL_LEN          uint32 = 0x3FFF0069
	VI_ATTR_CMDR_LA                     uint32 = 0x3FFF006B
	VI_ATTR_VXI_DEV_CLASS               uint32 = 0x3FFF006C
	VI_ATTR_MAINFRAME_LA                uint32 = 0x3FFF0070
	VI_ATTR_MANF_NAME                   uint32 = 0xBFFF0072
	VI_ATTR_MODEL_NAME                  uint32 = 0xBFFF0077
	VI_ATTR_VXI_VME_INTR_STATUS         uint32 = 0x3FFF008B
	VI_ATTR_VXI_TRIG_STATUS             uint32 = 0x3FFF008D
	VI_ATTR_VXI_VME_SYSFAIL_STATE       uint32 = 0x3FFF0094
	VI_ATTR_WIN_BASE_ADDR_32            uint32 = 0x3FFF0098
	VI_ATTR_WIN_SIZE_32                 uint32 = 0x3FFF009A
	VI_ATTR_ASRL_AVAIL_NUM              uint32 = 0x3FFF00AC
	VI_ATTR_MEM_BASE_32                 uint32 = 0x3FFF00AD
	VI_ATTR_ASRL_CTS_STATE              uint32 = 0x3FFF00AE
	VI_ATTR_ASRL_DCD_STATE              uint32 = 0x3FFF00AF
	VI_ATTR_ASRL_DSR_STATE              uint32 = 0x3FFF00B1
	VI_ATTR_ASRL_DTR_STATE              uint32 = 0x3FFF00B2
	VI_ATTR_ASRL_END_IN                 uint32 = 0x3FFF00B3
	VI_ATTR_ASRL_END_OUT                uint32 = 0x3FFF00B4
	VI_ATTR_ASRL_REPLACE_CHAR           uint32 = 0x3FFF00BE
	VI_ATTR_ASRL_RI_STATE               uint32 = 0x3FFF00BF
	VI_ATTR_ASRL_RTS_STATE              uint32 = 0x3FFF00C0
	VI_ATTR_ASRL_XON_CHAR               uint32 = 0x3FFF00C1
	VI_ATTR_ASRL_XOFF_CHAR              uint32 = 0x3FFF00C2
	VI_ATTR_WIN_ACCESS                  uint32 = 0x3FFF00C3
	VI_ATTR_RM_SESSION                  uint32 = 0x3FFF00C4
	VI_ATTR_VXI_LA                      uint32 = 0x3FFF00D5
	VI_ATTR_MANF_ID                     uint32 = 0x3FFF00D9
	VI_ATTR_MEM_SIZE_32                 uint32 = 0x3FFF00DD
	VI_ATTR_MEM_SPACE                   uint32 = 0x3FFF00DE
	VI_ATTR_MODEL_CODE                  uint32 = 0x3FFF00DF
	VI_ATTR_SLOT                        uint32 = 0x3FFF00E8
	VI_ATTR_INTF_INST_NAME              uint32 = 0xBFFF00E9
	VI_ATTR_IMMEDIATE_SERV              uint32 = 0x3FFF0100
	VI_ATTR_INTF_PARENT_NUM             uint32 = 0x3FFF0101
	VI_ATTR_RSRC_SPEC_VERSION           uint32 = 0x3FFF0170
	VI_ATTR_INTF_TYPE                   uint32 = 0x3FFF0171
	VI_ATTR_GPIB_PRIMARY_ADDR           uint32 = 0x3FFF0172
	VI_ATTR_GPIB_SECONDARY_ADDR         uint32 = 0x3FFF0173
	VI_ATTR_RSRC_MANF_NAME              uint32 = 0xBFFF0174
	VI_ATTR_RSRC_MANF_ID                uint32 = 0x3FFF0175
	VI_ATTR_INTF_NUM                    uint32 = 0x3FFF0176
	VI_ATTR_TRIG_ID                     uint32 = 0x3FFF0177
	VI_ATTR_GPIB_REN_STATE              uint32 = 0x3FFF0181
	VI_ATTR_GPIB_UNADDR_EN              uint32 = 0x3FFF0184
	VI_ATTR_DEV_STATUS_BYTE             uint32 = 0x3FFF0189
	VI_ATTR_FILE_APPEND_EN              uint32 = 0x3FFF0192
	VI_ATTR_VXI_TRIG_SUPPORT            uint32 = 0x3FFF0194
	VI_ATTR_TCPIP_ADDR                  uint32 = 0xBFFF0195
	VI_ATTR_TCPIP_HOSTNAME              uint32 = 0xBFFF0196
	VI_ATTR_TCPIP_PORT                  uint32 = 0x3FFF0197
	VI_ATTR_TCPIP_DEVICE_NAME           uint32 = 0xBFFF0199
	VI_ATTR_TCPIP_NODELAY               uint32 = 0x3FFF019A
	VI_ATTR_TCPIP_KEEPALIVE             uint32 = 0x3FFF019B
	VI_ATTR_4882_COMPLIANT              uint32 = 0x3FFF019F
	VI_ATTR_USB_SERIAL_NUM              uint32 = 0xBFFF01A0
	VI_ATTR_USB_INTFC_NUM               uint32 = 0x3FFF01A1
	VI_ATTR_USB_PROTOCOL                uint32 = 0x3FFF01A7
	VI_ATTR_USB_MAX_INTR_SIZE           uint32 = 0x3FFF01AF
	VI_ATTR_PXI_DEV_NUM                 uint32 = 0x3FFF0201
	VI_ATTR_PXI_FUNC_NUM                uint32 = 0x3FFF0202
	VI_ATTR_PXI_BUS_NUM                 uint32 = 0x3FFF0205
	VI_ATTR_PXI_CHASSIS                 uint32 = 0x3FFF0206
	VI_ATTR_PXI_SLOTPATH                uint32 = 0xBFFF0207
	VI_ATTR_PXI_SLOT_LBUS_LEFT          uint32 = 0x3FFF0208
	VI_ATTR_PXI_SLOT_LBUS_RIGHT         uint32 = 0x3FFF0209
	VI_ATTR_PXI_TRIG_BUS                uint32 = 0x3FFF020A
	VI_ATTR_PXI_STAR_TRIG_BUS           uint32 = 0x3FFF020B
	VI_ATTR_PXI_STAR_TRIG_LINE          uint32 = 0x3FFF020C
	VI_ATTR_PXI_SRC_TRIG_BUS            uint32 = 0x3FFF020D
	VI_ATTR_PXI_DEST_TRIG_BUS           uint32 = 0x3FFF020E
	VI_ATTR_PXI_MEM_TYPE_BAR0           uint32 = 0x3FFF0211
	VI_ATTR_PXI_MEM_TYPE_BAR1           uint32 = 0x3FFF0212
	VI_ATTR_PXI_MEM_TYPE_BAR2           uint32 = 0x3FFF0213
	VI_ATTR_PXI_MEM_TYPE_BAR3           uint32 = 0x3FFF0214
	VI_ATTR_PXI_MEM_TYPE_BAR4           uint32 = 0x3FFF0215
	VI_ATTR_PXI_MEM_TYPE_BAR5           uint32 = 0x3FFF0216
	VI_ATTR_PXI_MEM_BASE_BAR0_32        uint32 = 0x3FFF0221
	VI_ATTR_PXI_MEM_BASE_BAR1_32        uint32 = 0x3FFF0222
	VI_ATTR_PXI_MEM_BASE_BAR2_32        uint32 = 0x3FFF0223
	VI_ATTR_PXI_MEM_BASE_BAR3_32        uint32 = 0x3FFF0224
	VI_ATTR_PXI_MEM_BASE_BAR4_32        uint32 = 0x3FFF0225
	VI_ATTR_PXI_MEM_BASE_BAR5_32        uint32 = 0x3FFF0226
	VI_ATTR_PXI_MEM_BASE_BAR0_64        uint32 = 0x3FFF0228
	VI_ATTR_PXI_MEM_BASE_BAR1_64        uint32 = 0x3FFF0229
	VI_ATTR_PXI_MEM_BASE_BAR2_64        uint32 = 0x3FFF022A
	VI_ATTR_PXI_MEM_BASE_BAR3_64        uint32 = 0x3FFF022B
	VI_ATTR_PXI_MEM_BASE_BAR4_64        uint32 = 0x3FFF022C
	VI_ATTR_PXI_MEM_BASE_BAR5_64        uint32 = 0x3FFF022D
	VI_ATTR_PXI_MEM_SIZE_BAR0_32        uint32 = 0x3FFF0231
	VI_ATTR_PXI_MEM_SIZE_BAR1_32        uint32 = 0x3FFF0232
	VI_ATTR_PXI_MEM_SIZE_BAR2_32        uint32 = 0x3FFF0233
	VI_ATTR_PXI_MEM_SIZE_BAR3_32        uint32 = 0x3FFF0234
	VI_ATTR_PXI_MEM_SIZE_BAR4_32        uint32 = 0x3FFF0235
	VI_ATTR_PXI_MEM_SIZE_BAR5_32        uint32 = 0x3FFF0236
	VI_ATTR_PXI_MEM_SIZE_BAR0_64        uint32 = 0x3FFF0238
	VI_ATTR_PXI_MEM_SIZE_BAR1_64        uint32 = 0x3FFF0239
	VI_ATTR_PXI_MEM_SIZE_BAR2_64        uint32 = 0x3FFF023A
	VI_ATTR_PXI_MEM_SIZE_BAR3_64        uint32 = 0x3FFF023B
	VI_ATTR_PXI_MEM_SIZE_BAR4_64        uint32 = 0x3FFF023C
	VI_ATTR_PXI_MEM_SIZE_BAR5_64        uint32 = 0x3FFF023D
	VI_ATTR_PXI_IS_EXPRESS              uint32 = 0x3FFF0240
	VI_ATTR_PXI_SLOT_LWIDTH             uint32 = 0x3FFF0241
	VI_ATTR_PXI_MAX_LWIDTH              uint32 = 0x3FFF0242
	VI_ATTR_PXI_ACTUAL_LWIDTH           uint32 = 0x3FFF0243
	VI_ATTR_PXI_DSTAR_BUS               uint32 = 0x3FFF0244
	VI_ATTR_PXI_DSTAR_SET               uint32 = 0x3FFF0245
	VI_ATTR_PXI_ALLOW_WRITE_COMBINE     uint32 = 0x3FFF0246
	VI_ATTR_TCPIP_HISLIP_OVERLAP_EN     uint32 = 0x3FFF0300
	VI_ATTR_TCPIP_HISLIP_VERSION        uint32 = 0x3FFF0301
	VI_ATTR_TCPIP_HISLIP_MAX_MESSAGE_KB uint32 = 0x3FFF0302
	VI_ATTR_TCPIP_IS_HISLIP             uint32 = 0x3FFF0303

	VI_ATTR_JOB_ID              uint32 = 0x3FFF4006
	VI_ATTR_EVENT_TYPE          uint32 = 0x3FFF4010
	VI_ATTR_SIGP_STATUS_ID      uint32 = 0x3FFF4011
	VI_ATTR_RECV_TRIG_ID        uint32 = 0x3FFF4012
	VI_ATTR_INTR_STATUS_ID      uint32 = 0x3FFF4023
	VI_ATTR_STATUS              uint32 = 0x3FFF4025
	VI_ATTR_RET_COUNT_32        uint32 = 0x3FFF4026
	VI_ATTR_BUFFER              uint32 = 0x3FFF4027
	VI_ATTR_RECV_INTR_LEVEL     uint32 = 0x3FFF4041
	VI_ATTR_OPER_NAME           uint32 = 0xBFFF4042
	VI_ATTR_GPIB_RECV_CIC_STATE uint32 = 0x3FFF4193
	VI_ATTR_RECV_TCPIP_ADDR     uint32 = 0xBFFF4198
	VI_ATTR_USB_RECV_INTR_SIZE  uint32 = 0x3FFF41B0
	VI_ATTR_USB_RECV_INTR_DATA  uint32 = 0xBFFF41B1
	VI_ATTR_PXI_RECV_INTR_SEQ   uint32 = 0x3FFF4240
	VI_ATTR_PXI_RECV_INTR_DATA  uint32 = 0x3FFF4241

	VI_ATTR_WIN_BASE_ADDR_64 uint32 = 0x3FFF009B
	VI_ATTR_WIN_SIZE_64      uint32 = 0x3FFF009C
	VI_ATTR_MEM_BASE_64      uint32 = 0x3FFF00D0
	VI_ATTR_MEM_SIZE_64      uint32 = 0x3FFF00D1

	VI_ATTR_USER_DATA_64 uint32 = 0x3FFF000A
	VI_ATTR_RET_COUNT_64 uint32 = 0x3FFF4028
)

/*- Attributes platform dependent size ------------------------------------*/
func init() {
	if strings.Contains(runtime.GOARCH, "amd64") {
		const (
			VI_ATTR_USER_DATA uint32 = VI_ATTR_USER_DATA_64
			VI_ATTR_RET_COUNT uint32 = VI_ATTR_RET_COUNT_64

			VI_ATTR_WIN_BASE_ADDR     uint32 = VI_ATTR_WIN_BASE_ADDR_64
			VI_ATTR_WIN_SIZE          uint32 = VI_ATTR_WIN_SIZE_64
			VI_ATTR_MEM_BASE          uint32 = VI_ATTR_MEM_BASE_64
			VI_ATTR_MEM_SIZE          uint32 = VI_ATTR_MEM_SIZE_64
			VI_ATTR_PXI_MEM_BASE_BAR0 uint32 = VI_ATTR_PXI_MEM_BASE_BAR0_64
			VI_ATTR_PXI_MEM_BASE_BAR1 uint32 = VI_ATTR_PXI_MEM_BASE_BAR1_64
			VI_ATTR_PXI_MEM_BASE_BAR2 uint32 = VI_ATTR_PXI_MEM_BASE_BAR2_64
			VI_ATTR_PXI_MEM_BASE_BAR3 uint32 = VI_ATTR_PXI_MEM_BASE_BAR3_64
			VI_ATTR_PXI_MEM_BASE_BAR4 uint32 = VI_ATTR_PXI_MEM_BASE_BAR4_64
			VI_ATTR_PXI_MEM_BASE_BAR5 uint32 = VI_ATTR_PXI_MEM_BASE_BAR5_64
			VI_ATTR_PXI_MEM_SIZE_BAR0 uint32 = VI_ATTR_PXI_MEM_SIZE_BAR0_64
			VI_ATTR_PXI_MEM_SIZE_BAR1 uint32 = VI_ATTR_PXI_MEM_SIZE_BAR1_64
			VI_ATTR_PXI_MEM_SIZE_BAR2 uint32 = VI_ATTR_PXI_MEM_SIZE_BAR2_64
			VI_ATTR_PXI_MEM_SIZE_BAR3 uint32 = VI_ATTR_PXI_MEM_SIZE_BAR3_64
			VI_ATTR_PXI_MEM_SIZE_BAR4 uint32 = VI_ATTR_PXI_MEM_SIZE_BAR4_64
			VI_ATTR_PXI_MEM_SIZE_BAR5 uint32 = VI_ATTR_PXI_MEM_SIZE_BAR5_64
		)

	} else {
		const (
			VI_ATTR_USER_DATA uint32 = VI_ATTR_USER_DATA_32
			VI_ATTR_RET_COUNT uint32 = VI_ATTR_RET_COUNT_32

			VI_ATTR_WIN_BASE_ADDR     uint32 = VI_ATTR_WIN_BASE_ADDR_32
			VI_ATTR_WIN_SIZE          uint32 = VI_ATTR_WIN_SIZE_32
			VI_ATTR_MEM_BASE          uint32 = VI_ATTR_MEM_BASE_32
			VI_ATTR_MEM_SIZE          uint32 = VI_ATTR_MEM_SIZE_32
			VI_ATTR_PXI_MEM_BASE_BAR0 uint32 = VI_ATTR_PXI_MEM_BASE_BAR0_32
			VI_ATTR_PXI_MEM_BASE_BAR1 uint32 = VI_ATTR_PXI_MEM_BASE_BAR1_32
			VI_ATTR_PXI_MEM_BASE_BAR2 uint32 = VI_ATTR_PXI_MEM_BASE_BAR2_32
			VI_ATTR_PXI_MEM_BASE_BAR3 uint32 = VI_ATTR_PXI_MEM_BASE_BAR3_32
			VI_ATTR_PXI_MEM_BASE_BAR4 uint32 = VI_ATTR_PXI_MEM_BASE_BAR4_32
			VI_ATTR_PXI_MEM_BASE_BAR5 uint32 = VI_ATTR_PXI_MEM_BASE_BAR5_32
			VI_ATTR_PXI_MEM_SIZE_BAR0 uint32 = VI_ATTR_PXI_MEM_SIZE_BAR0_32
			VI_ATTR_PXI_MEM_SIZE_BAR1 uint32 = VI_ATTR_PXI_MEM_SIZE_BAR1_32
			VI_ATTR_PXI_MEM_SIZE_BAR2 uint32 = VI_ATTR_PXI_MEM_SIZE_BAR2_32
			VI_ATTR_PXI_MEM_SIZE_BAR3 uint32 = VI_ATTR_PXI_MEM_SIZE_BAR3_32
			VI_ATTR_PXI_MEM_SIZE_BAR4 uint32 = VI_ATTR_PXI_MEM_SIZE_BAR4_32
			VI_ATTR_PXI_MEM_SIZE_BAR5 uint32 = VI_ATTR_PXI_MEM_SIZE_BAR5_32
		)
	}
}

/*- Event Types -------------------------------------------------------------*/
const (
	VI_EVENT_IO_COMPLETION    uint32 = 0x3FFF2009
	VI_EVENT_TRIG             uint32 = 0xBFFF200A
	VI_EVENT_SERVICE_REQ      uint32 = 0x3FFF200B
	VI_EVENT_CLEAR            uint32 = 0x3FFF200D
	VI_EVENT_EXCEPTION        uint32 = 0xBFFF200E
	VI_EVENT_GPIB_CIC         uint32 = 0x3FFF2012
	VI_EVENT_GPIB_TALK        uint32 = 0x3FFF2013
	VI_EVENT_GPIB_LISTEN      uint32 = 0x3FFF2014
	VI_EVENT_VXI_VME_SYSFAIL  uint32 = 0x3FFF201D
	VI_EVENT_VXI_VME_SYSRESET uint32 = 0x3FFF201E
	VI_EVENT_VXI_SIGP         uint32 = 0x3FFF2020
	VI_EVENT_VXI_VME_INTR     uint32 = 0xBFFF2021
	VI_EVENT_PXI_INTR         uint32 = 0x3FFF2022
	VI_EVENT_TCPIP_CONNECT    uint32 = 0x3FFF2036
	VI_EVENT_USB_INTR         uint32 = 0x3FFF2037
	VI_ALL_ENABLED_EVENTS     uint32 = 0x3FFF7FFF
)

/*- Completion and Error Codes ----------------------------------------------*/
const (
	VI_SUCCESS_EVENT_EN         uint = 0x3FFF0002 /* 3FFF0002,  1073676290 */
	VI_SUCCESS_EVENT_DIS        uint = 0x3FFF0003 /* 3FFF0003,  1073676291 */
	VI_SUCCESS_QUEUE_EMPTY      uint = 0x3FFF0004 /* 3FFF0004,  1073676292 */
	VI_SUCCESS_TERM_CHAR        uint = 0x3FFF0005 /* 3FFF0005,  1073676293 */
	VI_SUCCESS_MAX_CNT          uint = 0x3FFF0006 /* 3FFF0006,  1073676294 */
	VI_SUCCESS_DEV_NPRESENT     uint = 0x3FFF007D /* 3FFF007D,  1073676413 */
	VI_SUCCESS_TRIG_MAPPED      uint = 0x3FFF007E /* 3FFF007E,  1073676414 */
	VI_SUCCESS_QUEUE_NEMPTY     uint = 0x3FFF0080 /* 3FFF0080,  1073676416 */
	VI_SUCCESS_NCHAIN           uint = 0x3FFF0098 /* 3FFF0098,  1073676440 */
	VI_SUCCESS_NESTED_SHARED    uint = 0x3FFF0099 /* 3FFF0099,  1073676441 */
	VI_SUCCESS_NESTED_EXCLUSIVE uint = 0x3FFF009A /* 3FFF009A,  1073676442 */
	VI_SUCCESS_SYNC             uint = 0x3FFF009B /* 3FFF009B,  1073676443 */

	VI_WARN_QUEUE_OVERFLOW uint = 0x3FFF000C /* 3FFF000C,  1073676300 */
	VI_WARN_CONFIG_NLOADED uint = 0x3FFF0077 /* 3FFF0077,  1073676407 */
	VI_WARN_NULL_OBJECT    uint = 0x3FFF0082 /* 3FFF0082,  1073676418 */
	VI_WARN_NSUP__STATE    uint = 0x3FFF0084 /* 3FFF0084,  1073676420 */
	VI_WARN_UNKNOWN_STATUS uint = 0x3FFF0085 /* 3FFF0085,  1073676421 */
	VI_WARN_NSUP_BUF       uint = 0x3FFF0088 /* 3FFF0088,  1073676424 */
	VI_WARN_EXT_FUNC_NIMPL uint = 0x3FFF00A9 /* 3FFF00A9,  1073676457 */

	_VI_ERROR uint = 0x80000000 /* -2147483647 -1 */

	VI_ERROR_SYSTEM_ERROR      = _VI_ERROR + 0x3FFF0000 /* BFFF0000, -1073807360 */
	VI_ERROR_INV_OBJECT        = _VI_ERROR + 0x3FFF000E /* BFFF000E, -1073807346 */
	VI_ERROR_RSRC_LOCKED       = _VI_ERROR + 0x3FFF000F /* BFFF000F, -1073807345 */
	VI_ERROR_INV_EXPR          = _VI_ERROR + 0x3FFF0010 /* BFFF0010, -1073807344 */
	VI_ERROR_RSRC_NFOUND       = _VI_ERROR + 0x3FFF0011 /* BFFF0011, -1073807343 */
	VI_ERROR_INV_RSRC_NAME     = _VI_ERROR + 0x3FFF0012 /* BFFF0012, -1073807342 */
	VI_ERROR_INV_ACC_MODE      = _VI_ERROR + 0x3FFF0013 /* BFFF0013, -1073807341 */
	VI_ERROR_TMO               = _VI_ERROR + 0x3FFF0015 /* BFFF0015, -1073807339 */
	VI_ERROR_CLOSING_FAILED    = _VI_ERROR + 0x3FFF0016 /* BFFF0016, -1073807338 */
	VI_ERROR_INV_DEGREE        = _VI_ERROR + 0x3FFF001B /* BFFF001B, -1073807333 */
	VI_ERROR_INV_JOB_ID        = _VI_ERROR + 0x3FFF001C /* BFFF001C, -1073807332 */
	VI_ERROR_NSUP_ATTR         = _VI_ERROR + 0x3FFF001D /* BFFF001D, -1073807331 */
	VI_ERROR_NSUP_ATTR_STATE   = _VI_ERROR + 0x3FFF001E /* BFFF001E, -1073807330 */
	VI_ERROR_ATTR_READONLY     = _VI_ERROR + 0x3FFF001F /* BFFF001F, -1073807329 */
	VI_ERROR_INV_LOCK_TYPE     = _VI_ERROR + 0x3FFF0020 /* BFFF0020, -1073807328 */
	VI_ERROR_INV_ACCESS_KEY    = _VI_ERROR + 0x3FFF0021 /* BFFF0021, -1073807327 */
	VI_ERROR_INV_EVENT         = _VI_ERROR + 0x3FFF0026 /* BFFF0026, -1073807322 */
	VI_ERROR_INV_MECH          = _VI_ERROR + 0x3FFF0027 /* BFFF0027, -1073807321 */
	VI_ERROR_HNDLR_NINSTALLED  = _VI_ERROR + 0x3FFF0028 /* BFFF0028, -1073807320 */
	VI_ERROR_INV_HNDLR_REF     = _VI_ERROR + 0x3FFF0029 /* BFFF0029, -1073807319 */
	VI_ERROR_INV_CONTEXT       = _VI_ERROR + 0x3FFF002A /* BFFF002A, -1073807318 */
	VI_ERROR_NENABLED          = _VI_ERROR + 0x3FFF002F /* BFFF002F, -1073807313 */
	VI_ERROR_ABORT             = _VI_ERROR + 0x3FFF0030 /* BFFF0030, -1073807312 */
	VI_ERROR_RAW_WR_PROT_VIOL  = _VI_ERROR + 0x3FFF0034 /* BFFF0034, -1073807308 */
	VI_ERROR_RAW_RD_PROT_VIOL  = _VI_ERROR + 0x3FFF0035 /* BFFF0035, -1073807307 */
	VI_ERROR_OUTP_PROT_VIOL    = _VI_ERROR + 0x3FFF0036 /* BFFF0036, -1073807306 */
	VI_ERROR_INP_PROT_VIOL     = _VI_ERROR + 0x3FFF0037 /* BFFF0037, -1073807305 */
	VI_ERROR_BERR              = _VI_ERROR + 0x3FFF0038 /* BFFF0038, -1073807304 */
	VI_ERROR_IN_PROGRESS       = _VI_ERROR + 0x3FFF0039 /* BFFF0039, -1073807303 */
	VI_ERROR_INV_SETUP         = _VI_ERROR + 0x3FFF003A /* BFFF003A, -1073807302 */
	VI_ERROR_QUEUE_ERROR       = _VI_ERROR + 0x3FFF003B /* BFFF003B, -1073807301 */
	VI_ERROR_ALLOC             = _VI_ERROR + 0x3FFF003C /* BFFF003C, -1073807300 */
	VI_ERROR_INV_MASK          = _VI_ERROR + 0x3FFF003D /* BFFF003D, -1073807299 */
	VI_ERROR_IO                = _VI_ERROR + 0x3FFF003E /* BFFF003E, -1073807298 */
	VI_ERROR_INV_FMT           = _VI_ERROR + 0x3FFF003F /* BFFF003F, -1073807297 */
	VI_ERROR_NSUP_FMT          = _VI_ERROR + 0x3FFF0041 /* BFFF0041, -1073807295 */
	VI_ERROR_LINE_IN_USE       = _VI_ERROR + 0x3FFF0042 /* BFFF0042, -1073807294 */
	VI_ERROR_NSUP_MODE         = _VI_ERROR + 0x3FFF0046 /* BFFF0046, -1073807290 */
	VI_ERROR_SRQ_NOCCURRED     = _VI_ERROR + 0x3FFF004A /* BFFF004A, -1073807286 */
	VI_ERROR_INV_SPACE         = _VI_ERROR + 0x3FFF004E /* BFFF004E, -1073807282 */
	VI_ERROR_INV_OFFSET        = _VI_ERROR + 0x3FFF0051 /* BFFF0051, -1073807279 */
	VI_ERROR_INV_WIDTH         = _VI_ERROR + 0x3FFF0052 /* BFFF0052, -1073807278 */
	VI_ERROR_NSUP_OFFSET       = _VI_ERROR + 0x3FFF0054 /* BFFF0054, -1073807276 */
	VI_ERROR_NSUP_VAR_WIDTH    = _VI_ERROR + 0x3FFF0055 /* BFFF0055, -1073807275 */
	VI_ERROR_WINDOW_NMAPPED    = _VI_ERROR + 0x3FFF0057 /* BFFF0057, -1073807273 */
	VI_ERROR_RESP_PENDING      = _VI_ERROR + 0x3FFF0059 /* BFFF0059, -1073807271 */
	VI_ERROR_NLISTENERS        = _VI_ERROR + 0x3FFF005F /* BFFF005F, -1073807265 */
	VI_ERROR_NCIC              = _VI_ERROR + 0x3FFF0060 /* BFFF0060, -1073807264 */
	VI_ERROR_NSYS_CNTLR        = _VI_ERROR + 0x3FFF0061 /* BFFF0061, -1073807263 */
	VI_ERROR_NSUP_OPER         = _VI_ERROR + 0x3FFF0067 /* BFFF0067, -1073807257 */
	VI_ERROR_INTR_PENDING      = _VI_ERROR + 0x3FFF0068 /* BFFF0068, -1073807256 */
	VI_ERROR_ASRL_PARITY       = _VI_ERROR + 0x3FFF006A /* BFFF006A, -1073807254 */
	VI_ERROR_ASRL_FRAMING      = _VI_ERROR + 0x3FFF006B /* BFFF006B, -1073807253 */
	VI_ERROR_ASRL_OVERRUN      = _VI_ERROR + 0x3FFF006C /* BFFF006C, -1073807252 */
	VI_ERROR_TRIG_NMAPPED      = _VI_ERROR + 0x3FFF006E /* BFFF006E, -1073807250 */
	VI_ERROR_NSUP_ALIGN_OFFSET = _VI_ERROR + 0x3FFF0070 /* BFFF0070, -1073807248 */
	VI_ERROR_USER_BUF          = _VI_ERROR + 0x3FFF0071 /* BFFF0071, -1073807247 */
	VI_ERROR_RSRC_BUSY         = _VI_ERROR + 0x3FFF0072 /* BFFF0072, -1073807246 */
	VI_ERROR_NSUP_WIDTH        = _VI_ERROR + 0x3FFF0076 /* BFFF0076, -1073807242 */
	VI_ERROR_INV_PARAMETER     = _VI_ERROR + 0x3FFF0078 /* BFFF0078, -1073807240 */
	VI_ERROR_INV_PROT          = _VI_ERROR + 0x3FFF0079 /* BFFF0079, -1073807239 */
	VI_ERROR_INV_SIZE          = _VI_ERROR + 0x3FFF007B /* BFFF007B, -1073807237 */
	VI_ERROR_WINDOW_MAPPED     = _VI_ERROR + 0x3FFF0080 /* BFFF0080, -1073807232 */
	VI_ERROR_NIMPL_OPER        = _VI_ERROR + 0x3FFF0081 /* BFFF0081, -1073807231 */
	VI_ERROR_INV_LENGTH        = _VI_ERROR + 0x3FFF0083 /* BFFF0083, -1073807229 */
	VI_ERROR_INV_MODE          = _VI_ERROR + 0x3FFF0091 /* BFFF0091, -1073807215 */
	VI_ERROR_SESN_NLOCKED      = _VI_ERROR + 0x3FFF009C /* BFFF009C, -1073807204 */
	VI_ERROR_MEM_NSHARED       = _VI_ERROR + 0x3FFF009D /* BFFF009D, -1073807203 */
	VI_ERROR_LIBRARY_NFOUND    = _VI_ERROR + 0x3FFF009E /* BFFF009E, -1073807202 */
	VI_ERROR_NSUP_INTR         = _VI_ERROR + 0x3FFF009F /* BFFF009F, -1073807201 */
	VI_ERROR_INV_LINE          = _VI_ERROR + 0x3FFF00A0 /* BFFF00A0, -1073807200 */
	VI_ERROR_FILE_ACCESS       = _VI_ERROR + 0x3FFF00A1 /* BFFF00A1, -1073807199 */
	VI_ERROR_FILE_IO           = _VI_ERROR + 0x3FFF00A2 /* BFFF00A2, -1073807198 */
	VI_ERROR_NSUP_LINE         = _VI_ERROR + 0x3FFF00A3 /* BFFF00A3, -1073807197 */
	VI_ERROR_NSUP_MECH         = _VI_ERROR + 0x3FFF00A4 /* BFFF00A4, -1073807196 */
	VI_ERROR_INTF_NUM_NCONFIG  = _VI_ERROR + 0x3FFF00A5 /* BFFF00A5, -1073807195 */
	VI_ERROR_CONN_LOST         = _VI_ERROR + 0x3FFF00A6 /* BFFF00A6, -1073807194 */
	VI_ERROR_NPERMISSION       = _VI_ERROR + 0x3FFF00A8 /* BFFF00A8, -1073807192 */
)

/*- Other VISA Definitions --------------------------------------------------*/
const (
	VI_FIND_BUFLEN uint = 256

	VI_INTF_GPIB     uint = 1
	VI_INTF_VXI      uint = 2
	VI_INTF_GPIB_VXI uint = 3
	VI_INTF_ASRL     uint = 4
	VI_INTF_PXI      uint = 5
	VI_INTF_TCPIP    uint = 6
	VI_INTF_USB      uint = 7

	VI_PROT_NORMAL        uint = 1
	VI_PROT_FDC           uint = 2
	VI_PROT_HS488         uint = 3
	VI_PROT_4882_STRS     uint = 4
	VI_PROT_USBTMC_VENDOR uint = 5

	VI_FDC_NORMAL uint = 1
	VI_FDC_STREAM uint = 2

	VI_LOCAL_SPACE     uint = 0
	VI_A16_SPACE       uint = 1
	VI_A24_SPACE       uint = 2
	VI_A32_SPACE       uint = 3
	VI_A64_SPACE       uint = 4
	VI_PXI_ALLOC_SPACE uint = 9
	VI_PXI_CFG_SPACE   uint = 10
	VI_PXI_BAR0_SPACE  uint = 11
	VI_PXI_BAR1_SPACE  uint = 12
	VI_PXI_BAR2_SPACE  uint = 13
	VI_PXI_BAR3_SPACE  uint = 14
	VI_PXI_BAR4_SPACE  uint = 15
	VI_PXI_BAR5_SPACE  uint = 16
	VI_OPAQUE_SPACE    uint = 0xFFFF

	VI_UNKNOWN_LA      int = -1
	VI_UNKNOWN_SLOT    int = -1
	VI_UNKNOWN_LEVEL   int = -1
	VI_UNKNOWN_CHASSIS int = -1

	VI_QUEUE         uint = 1
	VI_HNDLR         uint = 2
	VI_SUSPEND_HNDLR uint = 4
	VI_ALL_MECH      uint = 0xFFFF

	VI_ANY_HNDLR uint = 0

	VI_TRIG_ALL         int  = -2
	VI_TRIG_SW          int  = -1
	VI_TRIG_TTL0        uint = 0
	VI_TRIG_TTL1        uint = 1
	VI_TRIG_TTL2        uint = 2
	VI_TRIG_TTL3        uint = 3
	VI_TRIG_TTL4        uint = 4
	VI_TRIG_TTL5        uint = 5
	VI_TRIG_TTL6        uint = 6
	VI_TRIG_TTL7        uint = 7
	VI_TRIG_ECL0        uint = 8
	VI_TRIG_ECL1        uint = 9
	VI_TRIG_ECL2        uint = 10
	VI_TRIG_ECL3        uint = 11
	VI_TRIG_ECL4        uint = 12
	VI_TRIG_ECL5        uint = 13
	VI_TRIG_STAR_SLOT1  uint = 14
	VI_TRIG_STAR_SLOT2  uint = 15
	VI_TRIG_STAR_SLOT3  uint = 16
	VI_TRIG_STAR_SLOT4  uint = 17
	VI_TRIG_STAR_SLOT5  uint = 18
	VI_TRIG_STAR_SLOT6  uint = 19
	VI_TRIG_STAR_SLOT7  uint = 20
	VI_TRIG_STAR_SLOT8  uint = 21
	VI_TRIG_STAR_SLOT9  uint = 22
	VI_TRIG_STAR_SLOT10 uint = 23
	VI_TRIG_STAR_SLOT11 uint = 24
	VI_TRIG_STAR_SLOT12 uint = 25
	VI_TRIG_STAR_INSTR  uint = 26
	VI_TRIG_PANEL_IN    uint = 27
	VI_TRIG_PANEL_OUT   uint = 28
	VI_TRIG_STAR_VXI0   uint = 29
	VI_TRIG_STAR_VXI1   uint = 30
	VI_TRIG_STAR_VXI2   uint = 31

	VI_TRIG_PROT_DEFAULT   uint = 0
	VI_TRIG_PROT_ON        uint = 1
	VI_TRIG_PROT_OFF       uint = 2
	VI_TRIG_PROT_SYNC      uint = 5
	VI_TRIG_PROT_RESERVE   uint = 6
	VI_TRIG_PROT_UNRESERVE uint = 7

	VI_READ_BUF           uint = 1
	VI_WRITE_BUF          uint = 2
	VI_READ_BUF_DISCARD   uint = 4
	VI_WRITE_BUF_DISCARD  uint = 8
	VI_IO_IN_BUF          uint = 16
	VI_IO_OUT_BUF         uint = 32
	VI_IO_IN_BUF_DISCARD  uint = 64
	VI_IO_OUT_BUF_DISCARD uint = 128

	VI_FLUSH_ON_ACCESS uint = 1
	VI_FLUSH_WHEN_FULL uint = 2
	VI_FLUSH_DISABLE   uint = 3

	VI_NMAPPED    uint = 1
	VI_USE_OPERS  uint = 2
	VI_DEREF_ADDR uint = 3

	VI_TMO_IMMEDIATE uint = 0
	VI_TMO_INFINITE  uint = 0xFFFFFFFF

	VI_NO_LOCK        uint = 0
	VI_EXCLUSIVE_LOCK uint = 1
	VI_SHARED_LOCK    uint = 2
	VI_LOAD_CONFIG    uint = 4

	VI_NO_SEC_ADDR uint = 0xFFFF

	VI_ASRL_PAR_NONE  uint = 0
	VI_ASRL_PAR_ODD   uint = 1
	VI_ASRL_PAR_EVEN  uint = 2
	VI_ASRL_PAR_MARK  uint = 3
	VI_ASRL_PAR_SPACE uint = 4

	VI_ASRL_STOP_ONE  uint = 10
	VI_ASRL_STOP_ONE5 uint = 15
	VI_ASRL_STOP_TWO  uint = 20

	VI_ASRL_FLOW_NONE     uint = 0
	VI_ASRL_FLOW_XON_XOFF uint = 1
	VI_ASRL_FLOW_RTS_CTS  uint = 2
	VI_ASRL_FLOW_DTR_DSR  uint = 4

	VI_ASRL_END_NONE     uint = 0
	VI_ASRL_END_LAST_BIT uint = 1
	VI_ASRL_END_TERMCHAR uint = 2
	VI_ASRL_END_BREAK    uint = 3

	VI_STATE_ASSERTED   uint = 1
	VI_STATE_UNASSERTED uint = 0
	VI_STATE_UNKNOWN    int  = -1

	VI_BIG_ENDIAN    uint = 0
	VI_LITTLE_ENDIAN uint = 1

	VI_DATA_PRIV  uint = 0
	VI_DATA_NPRIV uint = 1
	VI_PROG_PRIV  uint = 2
	VI_PROG_NPRIV uint = 3
	VI_BLCK_PRIV  uint = 4
	VI_BLCK_NPRIV uint = 5
	VI_D64_PRIV   uint = 6
	VI_D64_NPRIV  uint = 7
	VI_D64_2EVME  uint = 8
	VI_D64_SST160 uint = 9
	VI_D64_SST267 uint = 10
	VI_D64_SST320 uint = 11

	VI_WIDTH_8  uint = 1
	VI_WIDTH_16 uint = 2
	VI_WIDTH_32 uint = 4
	VI_WIDTH_64 uint = 8

	VI_GPIB_REN_DEASSERT           uint = 0
	VI_GPIB_REN_ASSERT             uint = 1
	VI_GPIB_REN_DEASSERT_GTL       uint = 2
	VI_GPIB_REN_ASSERT_ADDRESS     uint = 3
	VI_GPIB_REN_ASSERT_LLO         uint = 4
	VI_GPIB_REN_ASSERT_ADDRESS_LLO uint = 5
	VI_GPIB_REN_ADDRESS_GTL        uint = 6

	VI_GPIB_ATN_DEASSERT               uint = 0
	VI_GPIB_ATN_ASSERT                 uint = 1
	VI_GPIB_ATN_DEASSERT_HANDSHAKEuint      = 2
	VI_GPIB_ATN_ASSERT_IMMEDIATE       uint = 3

	VI_GPIB_HS488_DISABLED uint = 0
	VI_GPIB_HS488_NIMPL    int  = -1

	VI_GPIB_UNADDRESSED uint = 0
	VI_GPIB_TALKER      uint = 1
	VI_GPIB_LISTENER    uint = 2

	VI_VXI_CMD16        uint = 0x0200
	VI_VXI_CMD16_RESP16 uint = 0x0202
	VI_VXI_RESP16       uint = 0x0002
	VI_VXI_CMD32        uint = 0x0400
	VI_VXI_CMD32_RESP16 uint = 0x0402
	VI_VXI_CMD32_RESP32 uint = 0x0404
	VI_VXI_RESP32       uint = 0x0004

	VI_ASSERT_SIGNAL       int  = -1
	VI_ASSERT_USE_ASSIGNED uint = 0
	VI_ASSERT_IRQ1         uint = 1
	VI_ASSERT_IRQ2         uint = 2
	VI_ASSERT_IRQ3         uint = 3
	VI_ASSERT_IRQ4         uint = 4
	VI_ASSERT_IRQ5         uint = 5
	VI_ASSERT_IRQ6         uint = 6
	VI_ASSERT_IRQ7         uint = 7

	VI_UTIL_ASSERT_SYSRESET  uint = 1
	VI_UTIL_ASSERT_SYSFAIL   uint = 2
	VI_UTIL_DEASSERT_SYSFAIL uint = 3

	VI_VXI_CLASS_MEMORY   uint = 0
	VI_VXI_CLASS_EXTENDED uint = 1
	VI_VXI_CLASS_MESSAGE  uint = 2
	VI_VXI_CLASS_REGISTER uint = 3
	VI_VXI_CLASS_OTHER    uint = 4

	VI_PXI_ADDR_NONE uint = 0
	VI_PXI_ADDR_MEM  uint = 1
	VI_PXI_ADDR_IO   uint = 2
	VI_PXI_ADDR_CFG  uint = 3

	VI_TRIG_UNKNOWN int = -1

	VI_PXI_LBUS_STAR_TRIG_BUS_0 uint = 1000
	VI_PXI_LBUS_STAR_TRIG_BUS_1 uint = 1001
	VI_PXI_LBUS_STAR_TRIG_BUS_2 uint = 1002
	VI_PXI_LBUS_STAR_TRIG_BUS_3 uint = 1003
	VI_PXI_LBUS_STAR_TRIG_BUS_4 uint = 1004
	VI_PXI_LBUS_STAR_TRIG_BUS_5 uint = 1005
	VI_PXI_LBUS_STAR_TRIG_BUS_6 uint = 1006
	VI_PXI_LBUS_STAR_TRIG_BUS_7 uint = 1007
	VI_PXI_LBUS_STAR_TRIG_BUS_8 uint = 1008
	VI_PXI_LBUS_STAR_TRIG_BUS_9 uint = 1009
	VI_PXI_STAR_TRIG_CONTROLLER uint = 1413
)

/*- Backward Compatibility Macros -----------------------------------------*/
var VI_ERROR_INV_SESSION = VI_ERROR_INV_OBJECT
var VI_INFINITE = VI_TMO_INFINITE
var VI_NORMAL = VI_PROT_NORMAL
var VI_FDC = VI_PROT_FDC
var VI_HS488 = VI_PROT_HS488
var VI_ASRL488 = VI_PROT_4882_STRS
var VI_ASRL_IN_BUF = VI_IO_IN_BUF
var VI_ASRL_OUT_BUF = VI_IO_OUT_BUF
var VI_ASRL_IN_BUF_DISCARD = VI_IO_IN_BUF_DISCARD
var VI_ASRL_OUT_BUF_DISCARD = VI_IO_OUT_BUF_DISCARD

func VI_VERSION_MAJOR(ver uint32) uint32 {
	return ver & 0xFFF00000 >> 20
}

func VI_VERSION_MINOR(ver uint32) uint32 {
	return ver & 0x000FFF00 >> 8
}

func VI_VERSION_SUBMINOR(ver uint32) uint32 {
	return ver & 0x000000FF
}

// This map context is from the guide of "NI-VISA Help"
var errorCode = map[uint]string{
	VI_ERROR_SYSTEM_ERROR:     " Unknown system error (miscellaneous error).",
	VI_ERROR_INV_OBJECT:       " The given session or object reference is invalid.",
	VI_ERROR_RSRC_LOCKED:      " Specified type of lock cannot be obtained or specified operation cannot be performed, because the resource is locked.",
	VI_ERROR_INV_EXPR:         " Invalid expression specified for search.",
	VI_ERROR_RSRC_NFOUND:      " Insufficient location information or the device or resource is not present in the system.",
	VI_ERROR_INV_RSRC_NAME:    " Invalid resource reference specified. Parsing error.",
	VI_ERROR_INV_ACC_MODE:     " Invalid access mode.",
	VI_ERROR_TMO:              " Timeout exfpired before operation completed.",
	VI_ERROR_CLOSING_FAILED:   " Unable to deallocate the previously allocated data structures corresponding to this session or object reference.",
	VI_ERROR_INV_DEGREE:       " Specified degree is invalid.",
	VI_ERROR_INV_JOB_ID:       " Specified job identifier is invalid.",
	VI_ERROR_NSUP_ATTR:        " The specified attribute is not defined or supported by the referenced session, event, or find list.",
	VI_ERROR_NSUP_ATTR_STATE:  " The specified state of the attribute is not valid, or is not supported as defined by the session, event, or find list.",
	VI_ERROR_ATTR_READONLY:    " The specified attribute is Read Only.",
	VI_ERROR_INV_LOCK_TYPE:    " The specified type of lock is not supported by this resource.",
	VI_ERROR_INV_ACCESS_KEY:   " The access key to the resource associated with this session is invalid.",
	VI_ERROR_INV_EVENT:        " Specified event type is not supported by the resource.",
	VI_ERROR_INV_MECH:         " Invalid mechanism specified.",
	VI_ERROR_HNDLR_NINSTALLED: " A handler is not currently installed for the specified event.",
	VI_ERROR_INV_HNDLR_REF:    " The given handler reference is invalid.",
	VI_ERROR_INV_CONTEXT:      " Specified event context is invalid.",
	//VI_ERROR_QUEUE_OVERFLOW :" The event queue for the specified type has overflowed (usually due to previous events not having been closed).",
	VI_ERROR_NENABLED:          " The session must be enabled for events of the specified type in order to receive them.",
	VI_ERROR_ABORT:             " The operation was aborted.",
	VI_ERROR_RAW_WR_PROT_VIOL:  " Violation of raw write protocol occurred during transfer.",
	VI_ERROR_RAW_RD_PROT_VIOL:  " Violation of raw read protocol occurred during transfer.",
	VI_ERROR_OUTP_PROT_VIOL:    " Device reported an output protocol error during transfer.",
	VI_ERROR_INP_PROT_VIOL:     " Device reported an input protocol error during transfer.",
	VI_ERROR_BERR:              " Bus error occurred during transfer.",
	VI_ERROR_IN_PROGRESS:       " Unable to queue the asynchronous operation because there is already an operation in progress.",
	VI_ERROR_INV_SETUP:         " Unable to start operation because setup is invalid (due to attributes being set to an inconsistent state).",
	VI_ERROR_QUEUE_ERROR:       " Unable to queue asynchronous operation (usually due to the I/O completion event not being enabled or insufficient space in the session's queue).",
	VI_ERROR_ALLOC:             " Insufficient system resources to perform necessary memory allocation.",
	VI_ERROR_INV_MASK:          " Invalid buffer mask specified.",
	VI_ERROR_IO:                " Could not perform operation because of I/O error.",
	VI_ERROR_INV_FMT:           " A format specifier in the format string is invalid.",
	VI_ERROR_NSUP_FMT:          " A format specifier in the format string is not supported.",
	VI_ERROR_LINE_IN_USE:       " The specified trigger line is currently in use.",
	VI_ERROR_NSUP_MODE:         " The specified mode is not supported by this VISA implementation.",
	VI_ERROR_SRQ_NOCCURRED:     " Service request has not been received for the session.",
	VI_ERROR_INV_SPACE:         " Invalid address space specified.",
	VI_ERROR_INV_OFFSET:        " Invalid offset specified.",
	VI_ERROR_INV_WIDTH:         " Invalid source or destination width specified.",
	VI_ERROR_NSUP_OFFSET:       " Specified offset is not accessible from this hardware.",
	VI_ERROR_NSUP_VAR_WIDTH:    " Cannot support source and destination widths that are different.",
	VI_ERROR_WINDOW_NMAPPED:    " The specified session is not currently mapped.",
	VI_ERROR_RESP_PENDING:      " A previous response is still pending, causing a multiple query error.",
	VI_ERROR_NLISTENERS:        " No Listeners condition is detected (both NRFD and NDAC are deasserted).",
	VI_ERROR_NCIC:              " The interface associated with this session is not currently the controller in charge.",
	VI_ERROR_NSYS_CNTLR:        " The interface associated with this session is not the system controller.",
	VI_ERROR_NSUP_OPER:         " The given session or object reference does not support this operation.",
	VI_ERROR_INTR_PENDING:      " An interrupt is still pending from a previous call.",
	VI_ERROR_ASRL_PARITY:       " A parity error occurred during transfer.",
	VI_ERROR_ASRL_FRAMING:      " A framing error occurred during transfer.",
	VI_ERROR_ASRL_OVERRUN:      " An overrun error occurred during transfer. A character was not read from the hardware before the next character arrived.",
	VI_ERROR_TRIG_NMAPPED:      " The path from trigSrc to trigDest is not currently mapped.",
	VI_ERROR_NSUP_ALIGN_OFFSET: " The specified offset is not properly aligned for the access width of the operation.",
	VI_ERROR_USER_BUF:          " A specified user buffer is not valid or cannot be accessed for the required size.",
	VI_ERROR_RSRC_BUSY:         " The resource is valid, but VISA cannot currently access it.",
	VI_ERROR_NSUP_WIDTH:        " Specified width is not supported by this hardware.",
	VI_ERROR_INV_PARAMETER:     " The value of some parameter—which parameter is not known—is invalid.",
	VI_ERROR_INV_PROT:          " The protocol specified is invalid.",
	VI_ERROR_INV_SIZE:          " Invalid size of window specified.",
	VI_ERROR_WINDOW_MAPPED:     " The specified session currently contains a mapped window.",
	VI_ERROR_NIMPL_OPER:        " The given operation is not implemented.",
	VI_ERROR_INV_LENGTH:        " Invalid length specified.",
	VI_ERROR_INV_MODE:          " The specified mode is invalid.",
	VI_ERROR_SESN_NLOCKED:      " The current session did not have any lock on the resource.",
	VI_ERROR_MEM_NSHARED:       " The device does not export any memory.",
	VI_ERROR_LIBRARY_NFOUND:    " A code library required by VISA could not be located or loaded.",
	VI_ERROR_NSUP_INTR:         " The interface cannot generate an interrupt on the requested level or with the requested statusID value.",
	VI_ERROR_INV_LINE:          " The value specified by the line parameter is invalid.",
	VI_ERROR_FILE_ACCESS:       " An error occurred while trying to open the specified file. Possible reasons include an invalid path or lack of access rights.",
	VI_ERROR_FILE_IO:           " An error occurred while performing I/O on the specified file.",
	VI_ERROR_NSUP_LINE:         " One of the specified lines (trigSrc or trigDest) is not supported by this VISA implementation, or the combination of lines is not a valid mapping.",
	VI_ERROR_NSUP_MECH:         " The specified mechanism is not supported for the given event type.",
	VI_ERROR_INTF_NUM_NCONFIG:  " The interface type is valid but the specified interface number is not configured.",
	VI_ERROR_CONN_LOST:         " The connection for the given session has been lost.",
	//VI_ERROR_MACHINE_NAVAIL :" The remote machine does not exist or is not accepting any connections.",
	VI_ERROR_NPERMISSION: " Access to the resource or remote machine is denied. This is due to lack of sufficient privileges for the current user or machine.",
}

var completedCode = map[uint]string{
	VI_SUCCESS:              "  Operation completed successfully.",
	VI_SUCCESS_EVENT_EN:     "  Specified event is already enabled for at least one of the specified mechanisms.",
	VI_SUCCESS_EVENT_DIS:    "  Specified event is already disabled for at least one of the specified mechanisms.",
	VI_SUCCESS_QUEUE_EMPTY:  "  Operation completed successfully, but queue was already empty.",
	VI_SUCCESS_TERM_CHAR:    "  The specified termination character was read.",
	VI_SUCCESS_MAX_CNT:      "  The number of bytes read is equal to the input count.",
	VI_WARN_QUEUE_OVERFLOW:  "  The event returned is valid. One or more events that occurred have not been raised because there was no room available on the queue at the time of their occurrence. This could happen because VI_ATTR_MAX_QUEUE_LENGTH is not set to a large enough value for your application and/or events are coming in faster than you are servicing them.",
	VI_WARN_CONFIG_NLOADED:  "  The specified configuration either does not exist or could not be loaded; using VISA-specified defaults.",
	VI_SUCCESS_DEV_NPRESENT: "  Session opened successfully, but the device at the specified address is not responding.",
	VI_SUCCESS_TRIG_MAPPED:  "  The path from trigSrc to trigDest is already mapped.",
	VI_SUCCESS_QUEUE_NEMPTY: "  Wait terminated successfully on receipt of an event notification. There is still at least one more event occurrence of the requested type(s) available for this session.",
	VI_WARN_NULL_OBJECT:     "  The specified object reference is uninitialized.",
	//VI_WARN_NSUP_ATTR_STATE :"  Although the specified state of the attribute is valid, it is not supported by this resource implementation.",
	VI_WARN_UNKNOWN_STATUS:      "  The status code passed to the operation could not be interpreted.",
	VI_WARN_NSUP_BUF:            "  The specified buffer is not supported.",
	VI_SUCCESS_NCHAIN:           "  Event handled successfully. Do not invoke any other handlers on this session for this event.",
	VI_SUCCESS_NESTED_SHARED:    "  Operation completed successfully, and this session has nested shared locks.",
	VI_SUCCESS_NESTED_EXCLUSIVE: "  Operation completed successfully, and this session has nested exclusive locks.",
	VI_SUCCESS_SYNC:             "  Asynchronous operation request was actually performed synchronously.",
	VI_WARN_EXT_FUNC_NIMPL:      "  The operation succeeded, but a lower level driver did not implement the extended functionality.",
}

