package payment

import (
	"strconv"

	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/constants"

	models "github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models"

	midtrans "github.com/veritrans/go-midtrans"
)

func GetPaymentURL(transaction models.Transaction, user models.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = constants.ServerKey
	midclient.ClientKey = constants.ClientKey
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReg := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			FName: user.Name,
			Email: user.Email,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReg)
	if err != nil {
		return "", err
	}
	return snapTokenResp.RedirectURL, nil
}
