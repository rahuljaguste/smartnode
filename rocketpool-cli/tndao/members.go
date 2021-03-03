package tndao

import (
    "fmt"

    "github.com/rocket-pool/rocketpool-go/utils/eth"
    "github.com/urfave/cli"

    "github.com/rocket-pool/smartnode/shared/services/rocketpool"
    "github.com/rocket-pool/smartnode/shared/utils/math"
)


func getMembers(c *cli.Context) error {

    // Get RP client
    rp, err := rocketpool.NewClientFromCtx(c)
    if err != nil { return err }
    defer rp.Close()

    // Get trusted node DAO members
    members, err := rp.TNDAOMembers()
    if err != nil {
        return err
    }

    // Print & return
    if len(members.Members) > 0 {
        fmt.Printf("The trusted node DAO has %d members:\n", len(members.Members))
        fmt.Println("")
    } else {
        fmt.Println("The trusted node DAO does not have any members yet.")
    }
    for _, member := range members.Members {
        fmt.Printf("--------------------\n")
        fmt.Printf("\n")
        fmt.Printf("Member ID:            %s\n", member.ID)
        fmt.Printf("Email address:        %s\n", member.Email)
        fmt.Printf("Joined at block:      %d\n", member.JoinedBlock)
        fmt.Printf("Last proposal block:  %d\n", member.LastProposalBlock)
        fmt.Printf("RPL bond amount:      %.6f\n", math.RoundDown(eth.WeiToEth(member.RPLBondAmount), 6))
        fmt.Printf("Unbonded minipools:   %d\n", member.UnbondedValidatorCount)
        fmt.Printf("\n")
    }
    return nil

}

