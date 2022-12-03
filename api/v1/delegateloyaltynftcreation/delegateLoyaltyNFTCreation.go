package delegateloyaltynftcreation

import (
	"net/http"

	"github.com/0xZkPay/platform/config/smartcontract/rawtransaction"
	"github.com/0xZkPay/platform/generated/smartcontract/zkpay"
	"github.com/0xZkPay/platform/util/pkg/httphelper"
	"github.com/0xZkPay/platform/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/delegateLoyaltyNFTCreation")
	{
		g.POST("", deletegateLoyaltyNFTCreation)
	}
}

func deletegateLoyaltyNFTCreation(c *gin.Context) {
	var request DelegateLoyaltyNFTCreationRequest
	err := c.BindJSON(&request)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "payload is invalid")
		return
	}

	creatorAddr := common.HexToAddress(request.CreatorAddress)
	zkPayABI := zkpay.ZkpayABI

	tx, err := rawtransaction.SendRawTransaction(zkPayABI, "delegateLoyaltyNFTCreation", creatorAddr, request.MetaDataHash)

	if err != nil {
		httphelper.NewInternalServerError(c, "", "failed to call %v of %v, error: %v", "delegateLoyaltyNFTCreation", "ZkPay", err.Error())
		return
	}
	transactionHash := tx.Hash().String()
	payload := DelegateLoyaltyNFTCreationPayload{
		TransactionHash: transactionHash,
	}
	logwrapper.Infof("Transaction Hash is %v", transactionHash)
	httphelper.SuccessResponse(c, "Transaction Sent, NFT Will Be Minted Soon", payload)
}
