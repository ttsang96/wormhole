package guardiand

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"vebridge/node/pkg/vaa"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/prototext"

	nodev1 "vebridge/node/proto/node/v1"
)

var AdminClientGovernanceVAAVerifyCmd = &cobra.Command{
	Use:   "governance-vaa-verify [FILENAME]",
	Short: "Verify governance vaa in prototxt format (offline)",
	Run:   runGovernanceVAAVerify,
	Args:  cobra.ExactArgs(1),
}

func runGovernanceVAAVerify(cmd *cobra.Command, args []string) {
	path := args[0]

	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	var msg nodev1.InjectGovernanceVAARequest
	err = prototext.Unmarshal(b, &msg)
	if err != nil {
		log.Fatalf("failed to deserialize: %v", err)
	}

	for _, message := range msg.Messages {
		var (
			v *vaa.VAA
		)
		switch payload := message.Payload.(type) {
		case *nodev1.GovernanceMessage_GuardianSet:
			v, err = adminGuardianSetUpdateToVAA(payload.GuardianSet, msg.CurrentSetIndex, message.Timestamp)
		case *nodev1.GovernanceMessage_ContractUpgrade:
			v, err = adminContractUpgradeToVAA(payload.ContractUpgrade, msg.CurrentSetIndex, message.Timestamp)
		}
		if err != nil {
			log.Fatalf("invalid update: %v", err)
		}

		digest, err := v.SigningMsg()
		if err != nil {
			panic(err)
		}

		log.Printf("Serialized: %v", hex.EncodeToString(b))

		log.Printf("VAA with digest %s: %+v", digest.Hex(), spew.Sdump(v))
	}
}
