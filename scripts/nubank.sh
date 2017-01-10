# Nubank
#
# export TOKEN= `./nubank.sh auth <cpf-sem-tracos> <senha>`
# ./nubank.sh transactions 2016-01-01
#

CT="Content-Type:application/json"
ORIG="Origin: https://conta.nubank.com.br"

function GET {
	curl -H "$CT" -H "$ORIG" -H "Authorization: Bearer $TOKEN" "$1"
}

function POST {
	curl -XPOST -H "$CT" -H "$ORIG" -d "$1" "$2"
}

## Authenticate
function auth {
	resp=$(curl -sL -H "$CT" -H "$ORIG" -H "X-Correlation-Id: WEB-APP.QAoe8" https://prod-s0-webapp-proxy.nubank.com.br/api/discovery)
	loginUrl=$(echo $resp | awk -F '\"' '{print $4}')
	resp=$(POST "{\"username\": \"$1\", \"password\": \"$2\", \"client_id\": \"other.conta\", \"client_secret\": \"yQPeLzoHuJzlMMSAjC-LgNUJdUecx8XO\", \"grant_type\": \"password\", \"nonce\": \"NOT-RANDOM-YET\"}" "$loginUrl")
	token=$(echo $resp | awk -F '\"' '{print $4}')
	echo $token
}

## Get transactions
function transactions {
	since="$1T21:55:55.496Z" #2015-01-01T21:55:55.496Z
	accountId="54c95e6b-f1d1-4af1-b8dd-2580c541bd92"
	GET "https://prod-s0-feed.nubank.com.br/api/accounts/$accountId/transactions?since=$since&transactions-version=v1.5"
}

case "$1" in
	auth) auth $2 $3 ;;
	transactions) transactions $2 ;;
	*) echo "Usage: export TOKEN=`$0 auth username password`" ;;
esac
