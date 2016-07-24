# Nubank

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
	resp=$(POST "{\"username\": \"$1\", \"password\": \"$2\", \"client_id\": \"other.legacy\", \"client_secret\": \"1iHY2WHAXj25GFSHTx9lyaTYnb4uB-v6\", \"grant_type\": \"password\", \"nonce\": \"NOT-RANDOM-YET\"}" "https://prod-auth.nubank.com.br/api/token")
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
