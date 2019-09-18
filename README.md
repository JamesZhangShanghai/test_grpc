This is a tool to simulate the cnBR to set up a GRPC connection with OPSHub and then send ALLCMMGMT and RANGING_CM messages to OPSHub.

It can be used in automation test cases.

How to use:

cnbr_tm_simulator <opshub-clusetr-ip> <cnbr-cluster-ip-start> <num-of-ips> <num-of-msgs> <proto-name>
  
example:
./test_grpc 10.124.210.120 192.168.235.134 1 1 RANGING_CM

./test_grpc 10.124.210.120 192.168.235.135 1 1 RANGING_CM_V2
