package main

import (
	"./chan_service"
	"./cmmgmt"
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

func main() {
	var req chan_service.Request
	var req_usresil chan_service.Request
	var cluster_ip_int int
	var num_of_ips int
	var num_of_msgs int

	if len(os.Args) < 5 {
		fmt.Println("Please input right parameter as: cnbr_sim <opshub-clusetr-ip> <cnbr-cluster-ip-start> <num-of-ips> <num-of-msgs>")
		return
	}

	cluster_ip_int = StringIpToInt(os.Args[2])
	num_of_ips, _ = strconv.Atoi(os.Args[3])
	num_of_msgs, _ = strconv.Atoi(os.Args[4])
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

	RemoteClient := chan_service.NewPostClient(RemoteConn)
	defer RemoteConn.Close()

	req.ProtoName = "ALLCMMGMT"
	req.UnixTime = time.Now().UnixNano()
	req.Namespace = "cmts-data"
	req.Expiration = 30
	req.IsDelta = true
	allCmsTm := cmmgmt.AllCmStatesTmMsg{}
	allCmsTm.AllCmStatesTm = append(allCmsTm.AllCmStatesTm, &CmStateTm)
	bytes, err1 := proto.Marshal(&allCmsTm)
	if err1 != nil {
		fmt.Println("Proto ALLCMMGMT marshal error - ", err)
	}
	req.TmData = bytes

	req_usresil.ProtoName = "USRESIL_STATS"
	req_usresil.UnixTime = time.Now().UnixNano()
	req_usresil.Namespace = "cmts-data"
	req_usresil.Expiration = 30
	req_usresil.IsDelta = true
	UsResilStatsList := usresiltm.UsResilStatsList{}
	UsResilStatsList.Stats = append(UsResilStatsList.Stats, &UsResilStatsMsg)
	bytes_usresil, err1 := proto.Marshal(&UsResilStatsList)
	if err1 != nil {
		fmt.Println("Proto USRESIL_STATS marshal error - ", err)
	}
	req_usresil.TmData = bytes_usresil

	loop_end := false

	msg_cnt := num_of_msgs
	for {
		for i := 0; i < num_of_ips; i++ {
			req.ClusterIp = IpIntToString(cluster_ip_int + i)
			_, err = RemoteClient.StreamTmData(context.Background(), &req)
			if err != nil {
				fmt.Println("GRPC Connection not available: ", err)
			}

			req_usresil.ClusterIp = IpIntToString(cluster_ip_int + i)
			_, err = RemoteClient.StreamTmData(context.Background(), &req_usresil)
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
