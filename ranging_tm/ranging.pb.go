// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ranging.proto

/*
Package ranging_tm is a generated protocol buffer package.

It is generated from these files:
	ranging.proto

It has these top-level messages:
	Ranging_Tm
	Ranging_Cm_Stats
	Ranging_Cm_Stats_List
	Ranging_Md_Stats
	Ranging_Md_Stats_List
*/
package ranging_tm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Ranging_Tm struct {
	TimingAdjust    int32  `protobuf:"varint,1,opt,name=Timing_adjust,json=timingAdjust" json:"Timing_adjust,omitempty"`
	PowerAdjust     int32  `protobuf:"varint,2,opt,name=Power_adjust,json=powerAdjust" json:"Power_adjust,omitempty"`
	FrequencyAdjust int32  `protobuf:"varint,3,opt,name=Frequency_adjust,json=frequencyAdjust" json:"Frequency_adjust,omitempty"`
	SgId            uint32 `protobuf:"varint,4,opt,name=Sg_id,json=sgId" json:"Sg_id,omitempty"`
	MdId            uint32 `protobuf:"varint,5,opt,name=Md_id,json=mdId" json:"Md_id,omitempty"`
	UcId            uint32 `protobuf:"varint,6,opt,name=Uc_id,json=ucId" json:"Uc_id,omitempty"`
	CmMac           string `protobuf:"bytes,7,opt,name=Cm_mac,json=cmMac" json:"Cm_mac,omitempty"`
}

func (m *Ranging_Tm) Reset()                    { *m = Ranging_Tm{} }
func (m *Ranging_Tm) String() string            { return proto.CompactTextString(m) }
func (*Ranging_Tm) ProtoMessage()               {}
func (*Ranging_Tm) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ranging_Tm) GetTimingAdjust() int32 {
	if m != nil {
		return m.TimingAdjust
	}
	return 0
}

func (m *Ranging_Tm) GetPowerAdjust() int32 {
	if m != nil {
		return m.PowerAdjust
	}
	return 0
}

func (m *Ranging_Tm) GetFrequencyAdjust() int32 {
	if m != nil {
		return m.FrequencyAdjust
	}
	return 0
}

func (m *Ranging_Tm) GetSgId() uint32 {
	if m != nil {
		return m.SgId
	}
	return 0
}

func (m *Ranging_Tm) GetMdId() uint32 {
	if m != nil {
		return m.MdId
	}
	return 0
}

func (m *Ranging_Tm) GetUcId() uint32 {
	if m != nil {
		return m.UcId
	}
	return 0
}

func (m *Ranging_Tm) GetCmMac() string {
	if m != nil {
		return m.CmMac
	}
	return ""
}

type Ranging_Cm_Stats struct {
	SgId                 uint32                                 `protobuf:"varint,1,opt,name=Sg_id,json=sgId" json:"Sg_id,omitempty"`
	MdId                 uint32                                 `protobuf:"varint,2,opt,name=Md_id,json=mdId" json:"Md_id,omitempty"`
	CmMac                string                                 `protobuf:"bytes,3,opt,name=Cm_mac,json=cmMac" json:"Cm_mac,omitempty"`
	MinPwrLoad           uint32                                 `protobuf:"varint,4,opt,name=Min_pwr_load,json=minPwrLoad" json:"Min_pwr_load,omitempty"`
	CmtxStats            []*Ranging_Cm_Stats_Ranging_Cmtx_Stats `protobuf:"bytes,5,rep,name=Cmtx_stats,json=cmtxStats" json:"Cmtx_stats,omitempty"`
	SmExhaustedCount     uint32                                 `protobuf:"varint,6,opt,name=Sm_exhausted_count,json=smExhaustedCount" json:"Sm_exhausted_count,omitempty"`
	CmdpwrPwrMax         uint32                                 `protobuf:"varint,7,opt,name=Cmdpwr_pwr_max,json=cmdpwrPwrMax" json:"Cmdpwr_pwr_max,omitempty"`
	CmdpwrPwrHi          uint32                                 `protobuf:"varint,8,opt,name=Cmdpwr_pwr_hi,json=cmdpwrPwrHi" json:"Cmdpwr_pwr_hi,omitempty"`
	CmdpwrNeq            uint32                                 `protobuf:"varint,9,opt,name=Cmdpwr_neq,json=cmdpwrNeq" json:"Cmdpwr_neq,omitempty"`
	CmdpwrLoadMinSet     uint32                                 `protobuf:"varint,10,opt,name=Cmdpwr_load_min_set,json=cmdpwrLoadMinSet" json:"Cmdpwr_load_min_set,omitempty"`
	CmdpwrMinDrw         uint32                                 `protobuf:"varint,11,opt,name=Cmdpwr_min_drw,json=cmdpwrMinDrw" json:"Cmdpwr_min_drw,omitempty"`
	CmdpwrMaxDrw         uint32                                 `protobuf:"varint,12,opt,name=Cmdpwr_max_drw,json=cmdpwrMaxDrw" json:"Cmdpwr_max_drw,omitempty"`
	CmdpwrPwrMaxNorm     uint32                                 `protobuf:"varint,13,opt,name=Cmdpwr_pwr_max_norm,json=cmdpwrPwrMaxNorm" json:"Cmdpwr_pwr_max_norm,omitempty"`
	CmdpwrPwrHiNorm      uint32                                 `protobuf:"varint,14,opt,name=Cmdpwr_pwr_hi_norm,json=cmdpwrPwrHiNorm" json:"Cmdpwr_pwr_hi_norm,omitempty"`
	CmdpwrLoadMinSetNorm uint32                                 `protobuf:"varint,15,opt,name=Cmdpwr_load_min_set_norm,json=cmdpwrLoadMinSetNorm" json:"Cmdpwr_load_min_set_norm,omitempty"`
	CmdpwrMinDrwNorm     uint32                                 `protobuf:"varint,16,opt,name=Cmdpwr_min_drw_norm,json=cmdpwrMinDrwNorm" json:"Cmdpwr_min_drw_norm,omitempty"`
	CmdpwrMaxDrwNorm     uint32                                 `protobuf:"varint,17,opt,name=Cmdpwr_max_drw_norm,json=cmdpwrMaxDrwNorm" json:"Cmdpwr_max_drw_norm,omitempty"`
}

func (m *Ranging_Cm_Stats) Reset()                    { *m = Ranging_Cm_Stats{} }
func (m *Ranging_Cm_Stats) String() string            { return proto.CompactTextString(m) }
func (*Ranging_Cm_Stats) ProtoMessage()               {}
func (*Ranging_Cm_Stats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ranging_Cm_Stats) GetSgId() uint32 {
	if m != nil {
		return m.SgId
	}
	return 0
}

func (m *Ranging_Cm_Stats) GetMdId() uint32 {
	if m != nil {
		return m.MdId
	}
	return 0
}

func (m *Ranging_Cm_Stats) GetCmMac() string {
	if m != nil {
		return m.CmMac
	}
	return ""
}

func (m *Ranging_Cm_Stats) GetMinPwrLoad() uint32 {
	if m != nil {
		return m.MinPwrLoad
	}
	return 0
}

func (m *Ranging_Cm_Stats) GetCmtxStats() []*Ranging_Cm_Stats_Ranging_Cmtx_Stats {
	if m != nil {
		return m.CmtxStats
	}
	return nil
}

func (m *Ranging_Cm_Stats) GetSmExhaustedCount() uint32 {
	if m != nil {
		return m.SmExhaustedCount
	}
	return 0
}

func (m *Ranging_Cm_Stats) GetCmdpwrPwrMax() uint32 {
	if m != nil {
		return m.CmdpwrPwrMax
	}
	return 0
}

func (m *Ranging_Cm_Stats) GetCmdpwrPwrHi() uint32 {
	if m != nil {
		return m.CmdpwrPwrHi
	}
	return 0
}

func (m *Ranging_Cm_Stats) GetCmdpwrNeq() uint32 {
	if m != nil {
		return m.CmdpwrNeq
	}
	return 0
}

func (m *Ranging_Cm_Stats) GetCmdpwrLoadMinSet() uint32 {
	if m != nil {
		return m.CmdpwrLoadMinSet
	}
	return 0
}

func (m *Ranging_Cm_Stats) GetCmdpwrMinDrw() uint32 {
	if m != nil {
		return m.CmdpwrMinDrw
	}
	return 0
}

func (m *Ranging_Cm_Stats) GetCmdpwrMaxDrw() uint32 {
	if m != nil {
		return m.CmdpwrMaxDrw
	}
	return 0
}

func (m *Ranging_Cm_Stats) GetCmdpwrPwrMaxNorm() uint32 {
	if m != nil {
		return m.CmdpwrPwrMaxNorm
	}
	return 0
}

func (m *Ranging_Cm_Stats) GetCmdpwrPwrHiNorm() uint32 {
	if m != nil {
		return m.CmdpwrPwrHiNorm
	}
	return 0
}

func (m *Ranging_Cm_Stats) GetCmdpwrLoadMinSetNorm() uint32 {
	if m != nil {
		return m.CmdpwrLoadMinSetNorm
	}
	return 0
}

func (m *Ranging_Cm_Stats) GetCmdpwrMinDrwNorm() uint32 {
	if m != nil {
		return m.CmdpwrMinDrwNorm
	}
	return 0
}

func (m *Ranging_Cm_Stats) GetCmdpwrMaxDrwNorm() uint32 {
	if m != nil {
		return m.CmdpwrMaxDrwNorm
	}
	return 0
}

type Ranging_Cm_Stats_Ranging_Cmtx_Stats struct {
	UcId                 uint32 `protobuf:"varint,1,opt,name=Uc_id,json=ucId" json:"Uc_id,omitempty"`
	RngState             string `protobuf:"bytes,2,opt,name=Rng_state,json=rngState" json:"Rng_state,omitempty"`
	Snr                  uint32 `protobuf:"varint,3,opt,name=Snr,json=snr" json:"Snr,omitempty"`
	DataMer              uint32 `protobuf:"varint,4,opt,name=Data_mer,json=dataMer" json:"Data_mer,omitempty"`
	RxPower              int32  `protobuf:"varint,5,opt,name=Rx_power,json=rxPower" json:"Rx_power,omitempty"`
	PwrReportLevel       uint32 `protobuf:"varint,6,opt,name=Pwr_report_level,json=pwrReportLevel" json:"Pwr_report_level,omitempty"`
	PwrPeakLevel         uint32 `protobuf:"varint,7,opt,name=Pwr_peak_level,json=pwrPeakLevel" json:"Pwr_peak_level,omitempty"`
	PhyPeakLevel         uint32 `protobuf:"varint,8,opt,name=Phy_peak_level,json=phyPeakLevel" json:"Phy_peak_level,omitempty"`
	DynPwrMax            uint32 `protobuf:"varint,9,opt,name=Dyn_pwr_max,json=dynPwrMax" json:"Dyn_pwr_max,omitempty"`
	DynPwrMin            uint32 `protobuf:"varint,10,opt,name=Dyn_pwr_min,json=dynPwrMin" json:"Dyn_pwr_min,omitempty"`
	PwrMinLevel          uint32 `protobuf:"varint,11,opt,name=Pwr_min_level,json=pwrMinLevel" json:"Pwr_min_level,omitempty"`
	PwrLoad              uint32 `protobuf:"varint,12,opt,name=Pwr_load,json=pwrLoad" json:"Pwr_load,omitempty"`
	TimingOffset         int32  `protobuf:"varint,13,opt,name=Timing_offset,json=timingOffset" json:"Timing_offset,omitempty"`
	InitTimingOffset     int32  `protobuf:"varint,14,opt,name=Init_timing_offset,json=initTimingOffset" json:"Init_timing_offset,omitempty"`
	TimingAdjustRawAvg   int32  `protobuf:"varint,15,opt,name=Timing_adjust_raw_avg,json=timingAdjustRawAvg" json:"Timing_adjust_raw_avg,omitempty"`
	TimingAdjustRawLtAvg int32  `protobuf:"varint,16,opt,name=Timing_adjust_raw_lt_avg,json=timingAdjustRawLtAvg" json:"Timing_adjust_raw_lt_avg,omitempty"`
	TimingAdjustRawMin   int32  `protobuf:"varint,17,opt,name=Timing_adjust_raw_min,json=timingAdjustRawMin" json:"Timing_adjust_raw_min,omitempty"`
	TimingAdjustRawMax   int32  `protobuf:"varint,18,opt,name=Timing_adjust_raw_max,json=timingAdjustRawMax" json:"Timing_adjust_raw_max,omitempty"`
	GoodCoefCnt          uint32 `protobuf:"varint,19,opt,name=Good_coef_cnt,json=goodCoefCnt" json:"Good_coef_cnt,omitempty"`
	ScaleCoefCnt         uint32 `protobuf:"varint,20,opt,name=Scale_coef_cnt,json=scaleCoefCnt" json:"Scale_coef_cnt,omitempty"`
	ImpulseCoefCnt       uint32 `protobuf:"varint,21,opt,name=Impulse_coef_cnt,json=impulseCoefCnt" json:"Impulse_coef_cnt,omitempty"`
	PreEquCoeffLoadCnt   uint32 `protobuf:"varint,22,opt,name=Pre_equ_coeff_load_cnt,json=preEquCoeffLoadCnt" json:"Pre_equ_coeff_load_cnt,omitempty"`
	PhyOpMode            uint32 `protobuf:"varint,23,opt,name=Phy_op_mode,json=phyOpMode" json:"Phy_op_mode,omitempty"`
	InitRngInterval      uint32 `protobuf:"varint,24,opt,name=Init_rng_interval,json=initRngInterval" json:"Init_rng_interval,omitempty"`
	HitCount             uint32 `protobuf:"varint,25,opt,name=Hit_count,json=hitCount" json:"Hit_count,omitempty"`
	MisCount             uint32 `protobuf:"varint,26,opt,name=Mis_count,json=misCount" json:"Mis_count,omitempty"`
	ConMisCount          uint32 `protobuf:"varint,27,opt,name=Con_mis_count,json=conMisCount" json:"Con_mis_count,omitempty"`
	PowerAdj             int32  `protobuf:"varint,28,opt,name=Power_adj,json=powerAdj" json:"Power_adj,omitempty"`
	CmdpwrRepPwr         uint32 `protobuf:"varint,29,opt,name=Cmdpwr_rep_pwr,json=cmdpwrRepPwr" json:"Cmdpwr_rep_pwr,omitempty"`
	CmdpwrCmdPwr         uint32 `protobuf:"varint,30,opt,name=Cmdpwr_cmd_pwr,json=cmdpwrCmdPwr" json:"Cmdpwr_cmd_pwr,omitempty"`
	CmdpwrPwrLow         uint32 `protobuf:"varint,31,opt,name=Cmdpwr_pwr_low,json=cmdpwrPwrLow" json:"Cmdpwr_pwr_low,omitempty"`
	CmdpwrPwrLoad        uint32 `protobuf:"varint,32,opt,name=Cmdpwr_pwr_load,json=cmdpwrPwrLoad" json:"Cmdpwr_pwr_load,omitempty"`
	CmdpwrRepPwrNorm     uint32 `protobuf:"varint,33,opt,name=Cmdpwr_rep_pwr_norm,json=cmdpwrRepPwrNorm" json:"Cmdpwr_rep_pwr_norm,omitempty"`
	CmdpwrCmdPwrNorm     uint32 `protobuf:"varint,34,opt,name=Cmdpwr_cmd_pwr_norm,json=cmdpwrCmdPwrNorm" json:"Cmdpwr_cmd_pwr_norm,omitempty"`
	CmdpwrPwrLowNorm     uint32 `protobuf:"varint,35,opt,name=Cmdpwr_pwr_low_norm,json=cmdpwrPwrLowNorm" json:"Cmdpwr_pwr_low_norm,omitempty"`
	CmdpwrPwrLoadNorm    uint32 `protobuf:"varint,36,opt,name=Cmdpwr_pwr_load_norm,json=cmdpwrPwrLoadNorm" json:"Cmdpwr_pwr_load_norm,omitempty"`
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) Reset()         { *m = Ranging_Cm_Stats_Ranging_Cmtx_Stats{} }
func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) String() string { return proto.CompactTextString(m) }
func (*Ranging_Cm_Stats_Ranging_Cmtx_Stats) ProtoMessage()    {}
func (*Ranging_Cm_Stats_Ranging_Cmtx_Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetUcId() uint32 {
	if m != nil {
		return m.UcId
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetRngState() string {
	if m != nil {
		return m.RngState
	}
	return ""
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetSnr() uint32 {
	if m != nil {
		return m.Snr
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetDataMer() uint32 {
	if m != nil {
		return m.DataMer
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetRxPower() int32 {
	if m != nil {
		return m.RxPower
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetPwrReportLevel() uint32 {
	if m != nil {
		return m.PwrReportLevel
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetPwrPeakLevel() uint32 {
	if m != nil {
		return m.PwrPeakLevel
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetPhyPeakLevel() uint32 {
	if m != nil {
		return m.PhyPeakLevel
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetDynPwrMax() uint32 {
	if m != nil {
		return m.DynPwrMax
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetDynPwrMin() uint32 {
	if m != nil {
		return m.DynPwrMin
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetPwrMinLevel() uint32 {
	if m != nil {
		return m.PwrMinLevel
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetPwrLoad() uint32 {
	if m != nil {
		return m.PwrLoad
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetTimingOffset() int32 {
	if m != nil {
		return m.TimingOffset
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetInitTimingOffset() int32 {
	if m != nil {
		return m.InitTimingOffset
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetTimingAdjustRawAvg() int32 {
	if m != nil {
		return m.TimingAdjustRawAvg
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetTimingAdjustRawLtAvg() int32 {
	if m != nil {
		return m.TimingAdjustRawLtAvg
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetTimingAdjustRawMin() int32 {
	if m != nil {
		return m.TimingAdjustRawMin
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetTimingAdjustRawMax() int32 {
	if m != nil {
		return m.TimingAdjustRawMax
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetGoodCoefCnt() uint32 {
	if m != nil {
		return m.GoodCoefCnt
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetScaleCoefCnt() uint32 {
	if m != nil {
		return m.ScaleCoefCnt
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetImpulseCoefCnt() uint32 {
	if m != nil {
		return m.ImpulseCoefCnt
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetPreEquCoeffLoadCnt() uint32 {
	if m != nil {
		return m.PreEquCoeffLoadCnt
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetPhyOpMode() uint32 {
	if m != nil {
		return m.PhyOpMode
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetInitRngInterval() uint32 {
	if m != nil {
		return m.InitRngInterval
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetHitCount() uint32 {
	if m != nil {
		return m.HitCount
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetMisCount() uint32 {
	if m != nil {
		return m.MisCount
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetConMisCount() uint32 {
	if m != nil {
		return m.ConMisCount
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetPowerAdj() int32 {
	if m != nil {
		return m.PowerAdj
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetCmdpwrRepPwr() uint32 {
	if m != nil {
		return m.CmdpwrRepPwr
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetCmdpwrCmdPwr() uint32 {
	if m != nil {
		return m.CmdpwrCmdPwr
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetCmdpwrPwrLow() uint32 {
	if m != nil {
		return m.CmdpwrPwrLow
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetCmdpwrPwrLoad() uint32 {
	if m != nil {
		return m.CmdpwrPwrLoad
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetCmdpwrRepPwrNorm() uint32 {
	if m != nil {
		return m.CmdpwrRepPwrNorm
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetCmdpwrCmdPwrNorm() uint32 {
	if m != nil {
		return m.CmdpwrCmdPwrNorm
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetCmdpwrPwrLowNorm() uint32 {
	if m != nil {
		return m.CmdpwrPwrLowNorm
	}
	return 0
}

func (m *Ranging_Cm_Stats_Ranging_Cmtx_Stats) GetCmdpwrPwrLoadNorm() uint32 {
	if m != nil {
		return m.CmdpwrPwrLoadNorm
	}
	return 0
}

type Ranging_Cm_Stats_List struct {
	CmStats []*Ranging_Cm_Stats `protobuf:"bytes,1,rep,name=Cm_stats,json=cmStats" json:"Cm_stats,omitempty"`
}

func (m *Ranging_Cm_Stats_List) Reset()                    { *m = Ranging_Cm_Stats_List{} }
func (m *Ranging_Cm_Stats_List) String() string            { return proto.CompactTextString(m) }
func (*Ranging_Cm_Stats_List) ProtoMessage()               {}
func (*Ranging_Cm_Stats_List) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ranging_Cm_Stats_List) GetCmStats() []*Ranging_Cm_Stats {
	if m != nil {
		return m.CmStats
	}
	return nil
}

type Ranging_Md_Stats struct {
	SgId           uint32 `protobuf:"varint,1,opt,name=Sg_id,json=sgId" json:"Sg_id,omitempty"`
	MdId           uint32 `protobuf:"varint,2,opt,name=Md_id,json=mdId" json:"Md_id,omitempty"`
	TotalRngReqs   uint32 `protobuf:"varint,3,opt,name=Total_rng_reqs,json=totalRngReqs" json:"Total_rng_reqs,omitempty"`
	RngReqs        uint32 `protobuf:"varint,4,opt,name=Rng_reqs,json=rngReqs" json:"Rng_reqs,omitempty"`
	InitRngReqs    uint32 `protobuf:"varint,5,opt,name=Init_rng_reqs,json=initRngReqs" json:"Init_rng_reqs,omitempty"`
	BInitRngReqs   uint32 `protobuf:"varint,6,opt,name=BInit_rng_reqs,json=bInitRngReqs" json:"BInit_rng_reqs,omitempty"`
	InvalidRngReqs uint32 `protobuf:"varint,7,opt,name=Invalid_rng_reqs,json=invalidRngReqs" json:"Invalid_rng_reqs,omitempty"`
	RngAborts      uint32 `protobuf:"varint,8,opt,name=Rng_aborts,json=rngAborts" json:"Rng_aborts,omitempty"`
	RngErrors      uint32 `protobuf:"varint,9,opt,name=Rng_errors,json=rngErrors" json:"Rng_errors,omitempty"`
}

func (m *Ranging_Md_Stats) Reset()                    { *m = Ranging_Md_Stats{} }
func (m *Ranging_Md_Stats) String() string            { return proto.CompactTextString(m) }
func (*Ranging_Md_Stats) ProtoMessage()               {}
func (*Ranging_Md_Stats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Ranging_Md_Stats) GetSgId() uint32 {
	if m != nil {
		return m.SgId
	}
	return 0
}

func (m *Ranging_Md_Stats) GetMdId() uint32 {
	if m != nil {
		return m.MdId
	}
	return 0
}

func (m *Ranging_Md_Stats) GetTotalRngReqs() uint32 {
	if m != nil {
		return m.TotalRngReqs
	}
	return 0
}

func (m *Ranging_Md_Stats) GetRngReqs() uint32 {
	if m != nil {
		return m.RngReqs
	}
	return 0
}

func (m *Ranging_Md_Stats) GetInitRngReqs() uint32 {
	if m != nil {
		return m.InitRngReqs
	}
	return 0
}

func (m *Ranging_Md_Stats) GetBInitRngReqs() uint32 {
	if m != nil {
		return m.BInitRngReqs
	}
	return 0
}

func (m *Ranging_Md_Stats) GetInvalidRngReqs() uint32 {
	if m != nil {
		return m.InvalidRngReqs
	}
	return 0
}

func (m *Ranging_Md_Stats) GetRngAborts() uint32 {
	if m != nil {
		return m.RngAborts
	}
	return 0
}

func (m *Ranging_Md_Stats) GetRngErrors() uint32 {
	if m != nil {
		return m.RngErrors
	}
	return 0
}

type Ranging_Md_Stats_List struct {
	MdStats []*Ranging_Md_Stats `protobuf:"bytes,1,rep,name=Md_stats,json=mdStats" json:"Md_stats,omitempty"`
}

func (m *Ranging_Md_Stats_List) Reset()                    { *m = Ranging_Md_Stats_List{} }
func (m *Ranging_Md_Stats_List) String() string            { return proto.CompactTextString(m) }
func (*Ranging_Md_Stats_List) ProtoMessage()               {}
func (*Ranging_Md_Stats_List) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Ranging_Md_Stats_List) GetMdStats() []*Ranging_Md_Stats {
	if m != nil {
		return m.MdStats
	}
	return nil
}

func init() {
	proto.RegisterType((*Ranging_Tm)(nil), "ranging_tm.Ranging_Tm")
	proto.RegisterType((*Ranging_Cm_Stats)(nil), "ranging_tm.Ranging_Cm_Stats")
	proto.RegisterType((*Ranging_Cm_Stats_Ranging_Cmtx_Stats)(nil), "ranging_tm.Ranging_Cm_Stats.Ranging_Cmtx_Stats")
	proto.RegisterType((*Ranging_Cm_Stats_List)(nil), "ranging_tm.Ranging_Cm_Stats_List")
	proto.RegisterType((*Ranging_Md_Stats)(nil), "ranging_tm.Ranging_Md_Stats")
	proto.RegisterType((*Ranging_Md_Stats_List)(nil), "ranging_tm.Ranging_Md_Stats_List")
}

func init() { proto.RegisterFile("ranging.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0xdd, 0x6e, 0x22, 0x37,
	0x1f, 0xc6, 0xc5, 0x66, 0x59, 0xc0, 0x7c, 0x84, 0x4c, 0x92, 0x7d, 0x27, 0x9b, 0x4d, 0x5e, 0x96,
	0x8d, 0x2a, 0xfa, 0x95, 0x55, 0xb7, 0x52, 0x7b, 0x9c, 0x92, 0xb4, 0x8b, 0xc4, 0x64, 0xd1, 0x90,
	0x1e, 0x5b, 0xce, 0x8c, 0x19, 0xdc, 0xc5, 0xf6, 0xe0, 0x31, 0x19, 0xb8, 0xaa, 0x9e, 0xf7, 0x3e,
	0x7a, 0x1b, 0xbd, 0x86, 0xca, 0x1f, 0x03, 0x86, 0x45, 0xad, 0x7a, 0x10, 0x29, 0x3c, 0xfe, 0xfd,
	0xf1, 0xdf, 0xf6, 0xe3, 0xc7, 0x80, 0xa6, 0x40, 0x2c, 0x21, 0x2c, 0xb9, 0x4e, 0x05, 0x97, 0xdc,
	0x03, 0xf6, 0x23, 0x94, 0xb4, 0xfb, 0x67, 0x09, 0x80, 0xd0, 0x7e, 0x7c, 0xa0, 0xde, 0x5b, 0xd0,
	0x7c, 0x20, 0x54, 0x7d, 0x40, 0xf1, 0x6f, 0x8b, 0x4c, 0xfa, 0xa5, 0x4e, 0xa9, 0x57, 0x0e, 0x1b,
	0x52, 0x8b, 0x37, 0x5a, 0xf3, 0xde, 0x80, 0xc6, 0x88, 0xe7, 0x58, 0x14, 0xcc, 0x33, 0xcd, 0xd4,
	0x53, 0xa5, 0x59, 0xe4, 0x4b, 0xd0, 0xfe, 0x59, 0xe0, 0xf9, 0x02, 0xb3, 0x68, 0x55, 0x60, 0x07,
	0x1a, 0x3b, 0x9c, 0x14, 0xba, 0x45, 0x8f, 0x41, 0x79, 0x9c, 0x40, 0x12, 0xfb, 0xcf, 0x3b, 0xa5,
	0x5e, 0x33, 0x7c, 0x9e, 0x25, 0x83, 0x58, 0x89, 0x41, 0xac, 0xc4, 0xb2, 0x11, 0x69, 0x6c, 0xc4,
	0x5f, 0x23, 0x25, 0xbe, 0x30, 0xe2, 0x22, 0x1a, 0xc4, 0xde, 0x29, 0x78, 0xd1, 0xa7, 0x90, 0xa2,
	0xc8, 0xaf, 0x74, 0x4a, 0xbd, 0x5a, 0x58, 0x8e, 0x68, 0x80, 0xa2, 0xee, 0x1f, 0x6d, 0xd0, 0x2e,
	0xd6, 0xd5, 0xa7, 0x70, 0x2c, 0x91, 0xcc, 0x36, 0x53, 0x95, 0xf6, 0x4d, 0xf5, 0xcc, 0x99, 0x6a,
	0xf3, 0xad, 0x07, 0xce, 0xb7, 0x7a, 0x1d, 0xd0, 0x08, 0x08, 0x83, 0x69, 0x2e, 0xe0, 0x8c, 0xa3,
	0xa2, 0x65, 0x40, 0x09, 0x1b, 0xe5, 0x62, 0xc8, 0x51, 0xec, 0xdd, 0x03, 0xd0, 0xa7, 0x72, 0x09,
	0x33, 0x35, 0xa1, 0x5f, 0xee, 0x1c, 0xf4, 0xea, 0xef, 0xdf, 0x5d, 0x6f, 0x36, 0xfc, 0x7a, 0xb7,
	0x29, 0x47, 0x90, 0x4b, 0x23, 0x85, 0xb5, 0x88, 0xca, 0xa5, 0x69, 0xf9, 0x1b, 0xe0, 0x8d, 0x29,
	0xc4, 0xcb, 0x29, 0x5a, 0x64, 0x12, 0xc7, 0x30, 0xe2, 0x0b, 0x26, 0xed, 0x06, 0xb4, 0x33, 0x7a,
	0x57, 0x0c, 0xf4, 0x95, 0xee, 0x5d, 0x81, 0x56, 0x9f, 0xc6, 0xaa, 0x3d, 0xf5, 0x47, 0xd1, 0x52,
	0x6f, 0x4a, 0x33, 0x6c, 0x44, 0x5a, 0x1d, 0xe5, 0x22, 0x40, 0x4b, 0xaf, 0x0b, 0x9a, 0x0e, 0x35,
	0x25, 0x7e, 0x55, 0x43, 0xf5, 0x35, 0xf4, 0x81, 0x78, 0x17, 0x6a, 0x1d, 0x9a, 0x61, 0x78, 0xee,
	0xd7, 0x34, 0x50, 0x33, 0xc0, 0x3d, 0x9e, 0x7b, 0xdf, 0x82, 0x63, 0x3b, 0xac, 0xf6, 0x01, 0x52,
	0xc2, 0x60, 0x86, 0xa5, 0x0f, 0x4c, 0x5f, 0x86, 0x53, 0xfb, 0x11, 0x10, 0x36, 0xc6, 0x6e, 0x5f,
	0x8a, 0x8c, 0x45, 0xee, 0xd7, 0xdd, 0xbe, 0x02, 0xc2, 0x6e, 0x45, 0xee, 0x52, 0x68, 0xa9, 0xa9,
	0xc6, 0x16, 0x85, 0x96, 0x8a, 0xda, 0x4c, 0x5d, 0x90, 0x8c, 0x0b, 0xea, 0x37, 0xdd, 0xa9, 0xcd,
	0x42, 0xef, 0xb9, 0xa0, 0xde, 0xd7, 0xc0, 0xdb, 0x5a, 0xac, 0xa1, 0x5b, 0x9a, 0x3e, 0x74, 0x56,
	0xac, 0xe1, 0x1f, 0x80, 0xbf, 0x67, 0x59, 0xa6, 0xe4, 0x50, 0x97, 0x9c, 0xec, 0xae, 0x4d, 0xd7,
	0x6d, 0x7a, 0xb2, 0xeb, 0x33, 0x25, 0x6d, 0xb7, 0x27, 0xb3, 0xc8, 0x5d, 0xdc, 0x2c, 0xd4, 0xe0,
	0x47, 0x5b, 0xb8, 0x5e, 0xad, 0xc2, 0x5f, 0xfd, 0x05, 0x80, 0xf7, 0xb9, 0x4b, 0x36, 0xd7, 0xa1,
	0xe4, 0x5c, 0x87, 0x73, 0x50, 0x0b, 0x59, 0xa2, 0xed, 0x87, 0xb5, 0xa3, 0x6b, 0x61, 0x55, 0xb0,
	0x44, 0x55, 0x60, 0xaf, 0x0d, 0x0e, 0xc6, 0x4c, 0x68, 0x4b, 0x37, 0xc3, 0x83, 0x8c, 0x09, 0xef,
	0x0c, 0x54, 0x6f, 0x91, 0x44, 0x90, 0x62, 0x61, 0xcd, 0x5c, 0x89, 0x91, 0x44, 0x01, 0xd6, 0x43,
	0xe1, 0x12, 0xea, 0x4b, 0xad, 0x6f, 0x61, 0x39, 0xac, 0x88, 0xa5, 0xbe, 0xf7, 0x5e, 0x0f, 0xb4,
	0x47, 0xb9, 0x80, 0x02, 0xa7, 0x5c, 0x48, 0x38, 0xc3, 0x4f, 0x78, 0x66, 0x2d, 0xd9, 0x4a, 0x73,
	0x11, 0x6a, 0x79, 0xa8, 0x54, 0x75, 0xa4, 0x8a, 0x4c, 0x31, 0xfa, 0x64, 0x39, 0x6b, 0x48, 0xb5,
	0xed, 0x18, 0x7d, 0xda, 0x50, 0xd3, 0x95, 0x4b, 0x55, 0x2d, 0x35, 0x5d, 0x6d, 0xa8, 0x4b, 0x50,
	0xbf, 0x5d, 0xb1, 0xb5, 0xb3, 0xad, 0x27, 0xe3, 0x15, 0xb3, 0xb6, 0x76, 0xc7, 0x09, 0xb3, 0x5e,
	0x2c, 0xc6, 0x09, 0x53, 0xb6, 0x1f, 0xd9, 0x13, 0x32, 0x93, 0x18, 0x0f, 0xd6, 0xcd, 0xd9, 0x98,
	0x39, 0xce, 0x40, 0x75, 0x54, 0x5c, 0x6e, 0x63, 0xbe, 0x8a, 0x3d, 0x6d, 0x27, 0x1a, 0xf9, 0x64,
	0xa2, 0xcc, 0xde, 0x74, 0xa3, 0xf1, 0xa3, 0xd6, 0xd4, 0x75, 0x1d, 0x30, 0x22, 0xa1, 0xdc, 0x22,
	0x5b, 0x9a, 0x6c, 0x13, 0x46, 0xe4, 0x83, 0x4b, 0x7f, 0x07, 0x4e, 0xb7, 0xd2, 0x16, 0x0a, 0x94,
	0x43, 0xf4, 0x94, 0x68, 0xaf, 0x95, 0x43, 0xcf, 0x4d, 0xdd, 0x10, 0xe5, 0x37, 0x4f, 0x89, 0x72,
	0xe8, 0xe7, 0x25, 0x33, 0xa9, 0xab, 0xda, 0xba, 0xea, 0x64, 0xa7, 0x6a, 0x28, 0x55, 0xdd, 0xde,
	0xa9, 0xd4, 0x36, 0x1d, 0xed, 0x9d, 0x4a, 0xed, 0xd7, 0xfe, 0x12, 0xb4, 0xf4, 0xbd, 0xfd, 0x25,
	0x26, 0x59, 0x7e, 0xe1, 0x5c, 0xa5, 0x14, 0x9e, 0xc0, 0x88, 0x49, 0xff, 0xd8, 0x6c, 0x71, 0xc2,
	0x79, 0xdc, 0xe7, 0x78, 0xd2, 0x37, 0x19, 0x35, 0x8e, 0xd0, 0x0c, 0x6f, 0xa0, 0x13, 0x73, 0xd8,
	0x99, 0x52, 0x0b, 0xaa, 0x07, 0xda, 0x03, 0x9a, 0x2e, 0x66, 0x99, 0xc3, 0x9d, 0x1a, 0x8b, 0x11,
	0xa3, 0x17, 0xe4, 0x7b, 0xf0, 0x72, 0x24, 0x30, 0xc4, 0xf3, 0x85, 0x26, 0x27, 0xe6, 0xea, 0x2a,
	0xfe, 0xa5, 0xe6, 0xbd, 0x54, 0xe0, 0xbb, 0xf9, 0x42, 0xe1, 0x13, 0x75, 0x90, 0xaa, 0xe6, 0x12,
	0xd4, 0x95, 0xe1, 0x78, 0x0a, 0x29, 0x8f, 0xb1, 0xff, 0x3f, 0x63, 0x95, 0x74, 0xba, 0xfa, 0x98,
	0x06, 0x3c, 0xc6, 0xde, 0x57, 0xe0, 0x48, 0x1f, 0xa3, 0x60, 0x09, 0x24, 0x4c, 0x62, 0xf1, 0x84,
	0x66, 0xbe, 0x6f, 0x32, 0x43, 0x9d, 0x62, 0xc8, 0x92, 0x81, 0x95, 0xd5, 0x8d, 0xfb, 0x40, 0xa4,
	0x0d, 0xe6, 0x33, 0xcd, 0x54, 0xa7, 0x44, 0x9a, 0x40, 0x3e, 0x07, 0xb5, 0x80, 0x64, 0x76, 0xf0,
	0x95, 0x19, 0xa4, 0x24, 0x33, 0x83, 0x2a, 0x87, 0x39, 0x83, 0x74, 0x0d, 0x9c, 0xdb, 0x1c, 0xe6,
	0x2c, 0x28, 0x98, 0x73, 0x50, 0x5b, 0xbf, 0xb5, 0xfe, 0x6b, 0xbd, 0xf1, 0xd5, 0xe2, 0xa1, 0x75,
	0x02, 0x53, 0xe0, 0x54, 0x19, 0xdf, 0xbf, 0x70, 0x03, 0x33, 0xc4, 0xe9, 0x28, 0x17, 0x0e, 0x15,
	0xd1, 0x58, 0x53, 0x97, 0x2e, 0xd5, 0xa7, 0xf1, 0x36, 0x65, 0xe2, 0x2f, 0xf7, 0xff, 0xbf, 0xf3,
	0x74, 0x0c, 0x79, 0xee, 0x7d, 0x01, 0x0e, 0xb7, 0x28, 0x14, 0xfb, 0x1d, 0x8d, 0x35, 0x1d, 0x0c,
	0xc5, 0x4e, 0xc2, 0xd9, 0xce, 0x4c, 0xc2, 0xbd, 0x71, 0x13, 0xce, 0xb4, 0xb7, 0x13, 0x88, 0xb6,
	0x45, 0x83, 0x77, 0x5d, 0xdc, 0xf4, 0xb9, 0x83, 0xdb, 0x5e, 0x0d, 0xfe, 0x76, 0xe7, 0x09, 0x18,
	0x72, 0x13, 0xb7, 0xef, 0xc0, 0xc9, 0x4e, 0xd3, 0x86, 0xbf, 0xd2, 0xfc, 0xd1, 0x56, 0xe7, 0xaa,
	0xa0, 0x3b, 0x02, 0xa7, 0xbb, 0xcf, 0x34, 0x1c, 0x92, 0x4c, 0x7a, 0x3f, 0x82, 0x6a, 0x9f, 0xda,
	0xb7, 0xbd, 0xa4, 0xdf, 0xf6, 0xd7, 0xff, 0xf4, 0xb6, 0x87, 0x95, 0x88, 0xea, 0x7f, 0xba, 0xbf,
	0x3f, 0xdb, 0xfc, 0x1c, 0x09, 0xe2, 0xff, 0xfc, 0x73, 0xe4, 0x0a, 0xb4, 0x1e, 0xb8, 0x44, 0x33,
	0x6d, 0x48, 0x81, 0xe7, 0x99, 0xcd, 0xf0, 0x86, 0x54, 0x6a, 0xc8, 0x92, 0x10, 0xcf, 0x33, 0x9d,
	0xd8, 0xc5, 0xb8, 0x0d, 0x73, 0x61, 0x87, 0xba, 0xa0, 0xb9, 0x36, 0xb4, 0x1e, 0x37, 0xbf, 0xab,
	0xea, 0xd6, 0xcc, 0x9a, 0xb9, 0x02, 0xad, 0x9f, 0xb6, 0x21, 0x93, 0xe9, 0x8d, 0xc7, 0x81, 0x43,
	0xa9, 0x8b, 0xc9, 0x9e, 0xd0, 0x8c, 0xc4, 0x1b, 0xae, 0x62, 0x2f, 0xa6, 0xd1, 0x0b, 0xf2, 0x02,
	0x00, 0xd5, 0x0e, 0x7a, 0xe4, 0x42, 0x66, 0x36, 0xd1, 0x6b, 0x82, 0x25, 0x37, 0x5a, 0x28, 0x86,
	0xb1, 0x10, 0x5c, 0x64, 0x45, 0x9a, 0x0b, 0x96, 0xdc, 0x69, 0xc1, 0x3d, 0x83, 0x62, 0xc3, 0xd6,
	0x67, 0x10, 0xc4, 0xff, 0x7e, 0x06, 0x45, 0x51, 0x58, 0xa1, 0xb1, 0xfe, 0xe7, 0xf1, 0x85, 0xfe,
	0xf5, 0xfb, 0xfd, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x1f, 0x27, 0x35, 0x0e, 0x0b, 0x00,
	0x00,
}
