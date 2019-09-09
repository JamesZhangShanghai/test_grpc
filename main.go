package main

import (
	"./chan_service"
	"./cmmgmt"
	"./ranging_tm"
	"./ranging_tm_v2"
	"./ranging_tm_v3"
	"./comm_util_pb"
	"./usresiltm"
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"strconv"
	"strings"
	"time"
)

var CmStateTm = cmmgmt.CmStateTmMsg{
	CmState: &comm_util_pb.ModemStateMsg{ClientName: "TestProto",
		MacAddr:    "0011.2233.4455",
		SvcGrp:     0,
		MacDomain:  0,
		PrimSid:    1,
		PrimDcid:   1,
		InitUcid:   1,
		PrimDsSfid: 1,
		PrimUsSfid: 1,
		MacState:   25,
		Ipv4Addr:   "10.22.33.44",
		Ipv6Addr:   []string{},
		GenBitmask: 0,
		RcsDcids:   []uint32{1, 2},
		TcsUcids:   []uint32{1, 2},
	},
	IsDelete: false,
}

var Rang_CmtxState=ranging_tm.Ranging_Cm_Stats_Ranging_Cmtx_Stats {
	        UcId: 3,
		RngState :  "test",
		Snr: 38,
		DataMer: 36,
                RxPower: 6,
		PwrReportLevel: 8,
		PwrPeakLevel: 3,
		PhyPeakLevel: 4,
		DynPwrMax: 12,
		DynPwrMin: 12,
		PwrMinLevel: 5,
		PwrLoad: 27,
		TimingOffset: -7,
		InitTimingOffset: -1,
		TimingAdjustRawAvg: 3,
		TimingAdjustRawLtAvg: -5,
		TimingAdjustRawMin: -9,
		TimingAdjustRawMax: -2,
		GoodCoefCnt: 9,
		ScaleCoefCnt: 10,
		ImpulseCoefCnt: 11,
		PreEquCoeffLoadCnt: 12,
		PhyOpMode: 2,
		InitRngInterval: 25,
		HitCount: 150,
		MisCount: 7,
		ConMisCount:1,
		PowerAdj: -2,
		CmdpwrRepPwr: 1,
		CmdpwrCmdPwr: 1,
		CmdpwrPwrLow: 1,
		CmdpwrPwrLoad: 1,
		CmdpwrRepPwrNorm: 1,
		CmdpwrCmdPwrNorm: 1,
		CmdpwrPwrLowNorm: 1,
		CmdpwrPwrLoadNorm: 1,
	}

var Rang_CmtxState_V2=ranging_tm_v2.Ranging_Cm_Stats_Ranging_Cmtx_Stats {
	        UcId: 3,
		RngState :  "test",
		Snr: 38,
		DataMer: 36,
                RxPower: 6,
		PwrReportLevel: 8,
		PwrPeakLevel: 3,
		PhyPeakLevel: 4,
		DynPwrMax: 12,
		DynPwrMin: 12,
		PwrMinLevel: 5,
		PwrLoad: 27,
		TimingOffset: -7,
		InitTimingOffset: -1,
		TimingAdjustRawAvg: 3,
		TimingAdjustRawLtAvg: -5,
		TimingAdjustRawMin: -9,
		TimingAdjustRawMax: -2,
		GoodCoefCnt: 9,
		ScaleCoefCnt: 10,
		ImpulseCoefCnt: 11,
		PreEquCoeffLoadCnt: 12,
		PhyOpMode: 2,
		InitRngInterval: 25,
		HitCount: 150,
		MisCount: 7,
		ConMisCount:1,
		PowerAdj: -2,
		CmdpwrRepPwr: 1,
		CmdpwrCmdPwr: 1,
		CmdpwrPwrLow: 1,
		CmdpwrPwrLoad: 1,
		CmdpwrRepPwrNorm: 1,
		CmdpwrCmdPwrNorm: 1,
		CmdpwrPwrLowNorm: 1,
		ReserveForV2Cnbr: "test_v2",
	}


var Rang_CmtxState_V3=ranging_tm_v3.Ranging_Cm_Stats_Ranging_Cmtx_Stats {
	        UcId: 3,
		RngState :  "test",
		Snr: 38,
		DataMer: 36,
                RxPower: 6,
		PwrReportLevel: 8,
		PwrPeakLevel: 3,
		PhyPeakLevel: 4,
		DynPwrMax: 12,
		DynPwrMin: 12,
		PwrMinLevel: 5,
		PwrLoad: 27,
		TimingOffset: -7,
		InitTimingOffset: -1,
		TimingAdjustRawAvg: 3,
		TimingAdjustRawLtAvg: -5,
		TimingAdjustRawMin: -9,
		TimingAdjustRawMax: -2,
		GoodCoefCnt: 9,
		ScaleCoefCnt: 10,
		ImpulseCoefCnt: 11,
		PreEquCoeffLoadCnt: 12,
		PhyOpMode: 2,
		InitRngInterval: 25,
		HitCount: 150,
		MisCount: 7,
		ConMisCount:1,
		PowerAdj: -2,
		CmdpwrRepPwr: 1,
		CmdpwrCmdPwr: 1,
		CmdpwrPwrLow: 1,
		CmdpwrPwrLoad: 1,
		CmdpwrRepPwrNorm: 1,
		CmdpwrCmdPwrNorm: 1,
		CmdpwrPwrLowNorm: 1,
		ReserveForV2Cnbr: "test_v3",
		MisCountV3: 90,
	}

var RangingTm = ranging_tm.Ranging_Cm_Stats{
	SgId: 0,
	MdId: 0,
	CmMac: "AA:BB:CC:DD:11:FF",
	MinPwrLoad: 6,
	CmtxStats: [] *ranging_tm.Ranging_Cm_Stats_Ranging_Cmtx_Stats {},
	SmExhaustedCount: 9,
	CmdpwrPwrMax: 8,
	CmdpwrPwrHi: 7,
	CmdpwrNeq: 32,
	CmdpwrLoadMinSet: 3,
	CmdpwrMinDrw: 1,
	CmdpwrMaxDrw: 1,
	CmdpwrPwrMaxNorm: 10,
	CmdpwrPwrHiNorm: 9,
	CmdpwrLoadMinSetNorm: 5,
	CmdpwrMinDrwNorm: 1,
	CmdpwrMaxDrwNorm: 1,
}

var RangingTmV2 = ranging_tm_v2.Ranging_Cm_Stats{
	SgId: 0,
	MdId: 0,
	CmMac: "AA:BB:CC:DD:EE:FF",
	MinPwrLoad: 6,
	CmtxStats: [] *ranging_tm_v2.Ranging_Cm_Stats_Ranging_Cmtx_Stats {},
	SmExhaustedCount: 9,
	CmdpwrPwrMax: 8,
	CmdpwrPwrHi: 7,
	CmdpwrNeq: 32,
	CmdpwrLoadMinSet: 3,
	CmdpwrMinDrw: 1,
	CmdpwrMaxDrw: 1,
	CmdpwrPwrMaxNorm: 10,
	CmdpwrPwrHiNorm: 9,
	CmdpwrLoadMinSetNorm: 5,
	CmdpwrMinDrwNorm: 1,
	CmdpwrMaxDrwNorm: 1,
}

var RangingTmV3 = ranging_tm_v3.Ranging_Cm_Stats{
	SgId: 0,
	MdId: 0,
	CmMac: "3F:BB:CC:DD:EE:FF",
	MinPwrLoad: 6,
	CmtxStats: [] *ranging_tm_v3.Ranging_Cm_Stats_Ranging_Cmtx_Stats {},
	SmExhaustedCount: 9,
	CmdpwrPwrMax: 8,
	CmdpwrPwrHi: 7,
	CmdpwrNeq: 32,
	CmdpwrLoadMinSet: 3,
	CmdpwrMinDrw: 1,
	CmdpwrMaxDrw: 1,
	CmdpwrPwrMaxNorm: 10,
	CmdpwrPwrHiNorm: 9,
	CmdpwrLoadMinSetNorm: 5,
	CmdpwrMinDrwNorm: 1,
	CmdpwrMaxDrwNorm: 1,
}

var UsResilStatsMsg = usresiltm.UsResilStats{
	SgId:       0,
	MdId:       0,
	ResilCmCnt: 6,
	TotalCmCnt: 8,
}



func StringIpToInt(ipstring string) int {
	ipSegs := strings.Split(ipstring, ".")
	var ipInt int = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}

func IpIntToString(ipInt int) string {
	ipSegs := make([]string, 4)
	var len int = len(ipSegs)
	buffer := bytes.NewBufferString("")
	for i := 0; i < len; i++ {
		tempInt := ipInt & 0xFF
		ipSegs[len-i-1] = strconv.Itoa(tempInt)
		ipInt = ipInt >> 8
	}
	for i := 0; i < len; i++ {
		buffer.WriteString(ipSegs[i])
		if i < len-1 {
			buffer.WriteString(".")
		}
	}
	return buffer.String()
}

func fill_allcmmgmt(req *chan_service.Request) {
	req.ProtoName = "ALLCMMGMT"
	req.UnixTime = time.Now().UnixNano()
        req.Namespace = "cmts-data"
        req.Expiration = 30
        req.IsDelta = true
        allCmsTm := cmmgmt.AllCmStatesTmMsg{}
        allCmsTm.AllCmStatesTm = append(allCmsTm.AllCmStatesTm, &CmStateTm)
        bytes, err1 := proto.Marshal(&allCmsTm)
        if err1 != nil {
               fmt.Println("Proto ALLCMMGMT marshal error - ", err1)
        }
        req.TmData = bytes

}

func fill_ranging_cm(req *chan_service.Request) {
	req.ProtoName = "RANGING_CM"
	req.UnixTime = time.Now().UnixNano()
        req.Namespace = "cmts-data"
        req.Expiration = 30
        req.IsDelta = true
	RangingTm.CmtxStats = append(RangingTm.CmtxStats, &Rang_CmtxState)
        RangingCmsTm := ranging_tm.Ranging_Cm_Stats_List{}
        RangingCmsTm.CmStats = append(RangingCmsTm.CmStats, &RangingTm)
        bytes, err1 := proto.Marshal(&RangingCmsTm)
        if err1 != nil {
               fmt.Println("Proto RANGING_CM marshal error - ", err1)
        }
        req.TmData = bytes

}

func fill_ranging_cm_v2(req *chan_service.Request) {
	req.ProtoName = "RANGING_CM_V2"
	req.UnixTime = time.Now().UnixNano()
        req.Namespace = "cmts-data"
        req.Expiration = 30
        req.IsDelta = true
	RangingTmV2.CmtxStats = append(RangingTmV2.CmtxStats, &Rang_CmtxState_V2)
        RangingCmsTm := ranging_tm_v2.Ranging_Cm_Stats_List{}
        RangingCmsTm.CmStats = append(RangingCmsTm.CmStats, &RangingTmV2)
        bytes, err1 := proto.Marshal(&RangingCmsTm)
        if err1 != nil {
               fmt.Println("Proto RANGING_CM marshal error - ", err1)
        }
        req.TmData = bytes
}

func fill_ranging_cm_v3(req *chan_service.Request) {
	req.ProtoName = "RANGING_CM_V3"
	req.UnixTime = time.Now().UnixNano()
        req.Namespace = "cmts-data"
        req.Expiration = 30
        req.IsDelta = true
	RangingTmV3.CmtxStats = append(RangingTmV3.CmtxStats, &Rang_CmtxState_V3)
        RangingCmsTm := ranging_tm_v3.Ranging_Cm_Stats_List{}
        RangingCmsTm.CmStats = append(RangingCmsTm.CmStats, &RangingTmV3)
        bytes, err1 := proto.Marshal(&RangingCmsTm)
        if err1 != nil {
               fmt.Println("Proto RANGING_CM marshal error - ", err1)
        }
        req.TmData = bytes
}

func main() {
	var req chan_service.Request
	var cluster_ip_int int
	var num_of_ips int
	var num_of_msgs int
	var protobuf_name string

	if len(os.Args) < 6 {
		fmt.Println("Please input right parameter as: cnbr_sim <opshub-clusetr-ip> <cnbr-cluster-ip-start> <num-of-ips> <num-of-msgs> <pb-name>")
		return
	}

	cluster_ip_int = StringIpToInt(os.Args[2])
	num_of_ips, _ = strconv.Atoi(os.Args[3])
	num_of_msgs, _ = strconv.Atoi(os.Args[4])
	protobuf_name = os.Args[5]
	//fmt.Printf("opshub ip: %s, cnbr_ip_int= %d, num_of_ips=%d \n", os.Args[1], cluster_ip_int, num_of_ips)
	// Set up a connection to the server.
	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})

	endPoint := fmt.Sprintf("telemetry.%s.nip.io:443", os.Args[1])
	RemoteConn, err := grpc.Dial(endPoint, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("Could not connect to", endPoint, " error:", err)
	} else {
		fmt.Println("Connected to", endPoint, " error:", err)
	}
        fmt.Println("Protobuf name%s:", protobuf_name)
	RemoteClient := chan_service.NewPostClient(RemoteConn)
	defer RemoteConn.Close()
        switch protobuf_name {
	case "ALLCMMGMT":
            fill_allcmmgmt(&req)
	case "RANGING_CM":
            fill_ranging_cm(&req)
	case "RANGING_CM_V2":
            fill_ranging_cm_v2(&req)
	case "RANGING_CM_V3":
            fill_ranging_cm_v3(&req)
	}

	loop_end := false

	msg_cnt := num_of_msgs
	for {
		for i := 0; i < num_of_ips; i++ {
			req.ClusterIp = IpIntToString(cluster_ip_int + i)
			_, err = RemoteClient.StreamTmData(context.Background(), &req)
			if err != nil {
				fmt.Println("GRPC Connection not available: ", err)
			}

			if num_of_msgs != 0 {
				if msg_cnt > 0 {
					msg_cnt = msg_cnt - 1
				} else {
					loop_end = true
					break
				}
			} else {
				msg_cnt = msg_cnt + 1
			}

			fmt.Printf(".")
			if msg_cnt > 0 && (msg_cnt%16 == 0) {
				fmt.Printf("\n")
			}
		}

		if loop_end {
			break
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}
