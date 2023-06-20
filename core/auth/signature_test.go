package auth

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifySignature(t *testing.T) {
	signature := "sJg4l85B61ldoYcCKwICu9nb/OFjr2zEHW5+YpmK6Oi9fHikfzPs4AfxJOqiUXSz2smyRlqNDPMR4oIkfbbAtbrpm810CwSJ7gjdG8IPbRBM9GB82vBPYLm3Oni7l6F3168i1WhB/eLKZUOQlD9WDx6/NiywQ+uZLj2CHzXYC+LThenE/ejNu2Y/NZPiNZVe7oOoiB7aOacg2UOJPwArlXSZuCjfXLbd5eDAl0yaPbzgysqX0Ohcx9KcSkDGVQ5RqadCHqfcRzZDi5nx3GTjKW8BzP2yFecaeQnZl537wbDBEzhTAw9BNusv3Gur3KwsxGeOscV14fPG4aOf/Z+JDN8cJnVLg6Urnekqvs/UfmoPlZ7xtdlLQsh4tQD18EeQ9tN2yNXmqniabNTc2W/TS/VBCUFgbOdsLKG3eFbfP8uj/pqMDtGsA08rJ/4OxYs8Rg7ZiNGKbwUKFHaSvsCxBvvpINAianUuB0DzTyeV3PstnhmY9FXaBIh1+0eZsdQdAwF+bZ7HYTkrqTalgNppdP0tJfVLhyp8zd2/tJPBdQ78lzUuoYCxLn2E7UVX1KUPXO82N/aMZTlw17DT7Rjh85jAtuipTwq/bUsiQUIMwhSBrR/yMedmgHhoBPHG+XWkg1pp6e5gEtn2LgF5OTp/3MZW3JKpniIcR+/j4Ku+uu0="

	json := `{"customer":{"email":"zaher.leon@gmail.com","company_name":"","reference":"","phone":"","first_name":"Leon","last_name":"Zaher","city":"Zagreb","postcode":"10000","state":"Grad Zagreb","country":"Croatia","customer_account":null,"metadata":{}},"custom_fields":[{"name":"abcd","data_type":null,"value":"efgh"}],"product_features":[],"validity_period":null,"allow_grace_period":false,"grace_period":0,"license_type":"perpetual","license_key":"GVMX-U4JH-UBAK-6ECO","is_trial":false,"maintenance_period":null,"max_activations":1,"allow_unlimited_activations":false,"times_activated":1,"allow_overages":false,"max_overages":0,"prevent_vm":false,"is_floating_cloud":false,"is_floating":true,"floating_timeout":120,"transfer_count":8,"max_transfers":0,"start_date":null,"product_details":{"product_id":1586066454419254,"product_name":"core","short_code":"core","authorization_method":"license-key","metadata":{}},"can_borrow":true,"max_borrow_time":1,"device_id":1687283003863909,"id":1673449189923780,"metadata":{},"floating_users":100,"license_signature":"ubk1m9jD3nNr1tRWAWiG6pS1ckzQCO1VDvVikM1C4xDSc3yO5+CoYiU92uZkIg+nW8iSVXa1s3faE83mh6XYotyfeJIFuDyLbQU51MMhzXxq1WW72kAgyRZIwEZ850EsqsJ6pIyLmCBjwCrjaTmIhzEo39iIJO6kY1oiv7UReRrZO93y6j1rNGGhJVxQH4EbaHwKZXzzdVx3GQa5BiosGNdnRgEsBbTyySU4bwePwWN8Ir3ierHNlgA5JB8EVNJf9Xd7Y2eS42xy3h/0k7VhUKYgpc8iKtl1YG16SupXEtMA69HdAoNByUNTBXCULDgEeDJ0s3yJIOm7bkIoscSV5szcWPwVfE02PAvHZvJ2SdEde56t4ngc5XjpUk3Gv0WwgQaJPjxkiE4PIPp1Ks38NvYM1MWQXGama9PhUm7JwYE1cHbL5mcWni0Mpz9egsOzoVD3Ezm86k3bSZ4mY7YNLqDjc6p+F/R/xyFuQ/0hlrK0lDp8v4NeG65xKRzbP1zzk9ucicc6VR8lM10FRd1wmsCsZeP4DPmmYxmgqwnkaeceELaEigciowF3IBN5mchSBE/9J84bL8BRvizPLi7Ac5aiSyQibpUPPbUHbYhW+cie2mpvLVv9OEFWTjKoknRcux83CX0GxCInLaToJpSV7izoWYhqFTv+kIha5YeVor0="}`

	err := VerifySignatureV2(signature, json)
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.NoError(t, err)
}
