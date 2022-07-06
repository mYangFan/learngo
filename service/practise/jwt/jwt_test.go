package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAjAPIH38hVOxEi9LiRUURwYP2ADufPc5f2xajLnMXMwVC6H+C
tyeNAFZ2hRwRdlVZskKvuebSRpndAPyQXP2poREwlgP+J/4/ZTTomqW89g38vgtn
VDwthStv0lLQab307zhQx+FvXFdjMIPSJsGCwx5GKtTATJeoiyDuwI8p9utoL3xl
s1Z0lQ2pvyYOCbf64VDmhtmlUxl8n5iBPgMGoE8LqO9FfsY9jTzQqi0tdM6EVWkn
hUhnaM+YF9l3zTB2lYp5IVvchmz1RmIUumNqlwBVQFvnjFSToLQNq8oHDxcvfOCt
a8j5Lr/0V3eZ4SorQBs33qxxotWrVtNpe20usQIDAQABAoIBAEdeAhqz1UyRJZtT
wYvnWdaWqcSSn2eEku18i264sUGLVABoRjuPSFq8t6q/lNMJTDAFt8Z9Nf9QnArN
sdQEY/s6ZNaCcID4Il426r0w77FHjeVcsK2KAvXXHydImPzdG0T4IozPhtC9pkxr
VKn2hu1nuAJK8T3ZK8rG3Yt/Lbeyox0blqjQIjdC5QcoiVjA2OAm74Du65/0nsUc
0aRgL8PNhCNiiGrkxWUdSTcqtWhSlKcO4hz+twm9EbaFaXAskCFViaEj2X1pyQ4P
oYAcLe9i+VMd4nhdPvgwzti23f3vew8eRHjtPzsAVzmXClxgYHezhZnivh+VWolt
PbhoJqUCgYEAwPed3EydM447VbdniBiWpL69gzidT65M9sgiuOdcKFZNTEgpevzD
s27NyoBQoKoY2WPRAv7HTpj3HJvAQhK52kXcUKoFUvLjqIOmkHb+SEhlfjUYZ46n
32xpEyFTIvP/c2qMuRwJt6So53hGN3XlislNe/gPSOJiMoqAXPD24/sCgYEAucAm
6a7FH+Tq1X9WO+7cPrcgfci5556W7xRF9s0KDG5mFDDkuPhu0qmk34VD063fNEjU
7a2yZWfEOWXN6I1dDOeO/aUV6euMvq4cwzDAjUUqASeEeXekS2MMvDmKMKaOu+Ab
O2mFjGj+nAzL54n0b9iZE1qVV6J7p9SzzMKTTEMCgYEAobNd32t8F+XKJuwHn+4q
3mbP3BQnDsxaxjmzjwqhGVQodhHKYLtLKyNg455e8iNXq2OPlFj3nw9jk9+YXkXe
49/C3P18dKQAzgd0Hn74Wo4ALqBDkRPj7L+l1VgJmLKqj4br4XpzUiZzO3R4MqVi
8sf5XDbkaAj8jKmvIzLxUN0CgYEAs21jqVODYz2zgiwQ1q/y3Dn1Dsv4mD5fOe5x
bue9DykNPe/E4NBJ7QCmHKwUSsOn2k+IL1cb/kxwBanLbxouiiqbu3PeaSl5uS9i
5UGQnmXzH05W7yac851oTeGfFUOumNjwNUHHGUzKtV6/EJf4IJovs7xKim6P2Bzk
2aQSuRUCgYBgzFQn2cgSd7Y74ve27HFqTZFXL0Ij7ZSkcH9IG48vZ0GXZvZZodfB
y377Ek4aq5ngZEkf4Ve1zhuDNzMFnzElVVnHoveDZvSXDUb3EkmpausoDKUU2Syk
632sbEkU4mhrFzvLEwhqDaTsM7QvaFxgEBJGUpuqFdDMF2vST3sfTQ==
-----END RSA PRIVATE KEY-----`

func TestGenerateToken(t *testing.T)  {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Fatalf("cannot Parse private key: %v", err)
	}
	g := NewJWTTokenGen("coolcar/auth", key)
	g.nowFunc = func() time.Time {
		return time.Unix(1516239022, 0)
	}

	tkn, err := g.GenerateToken("1234567890", 2 * time.Hour)

	fmt.Println(g.Issue)

	if err != nil {
		t.Errorf("cannot generate token error: %v", err)
	}

	want := "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiMTIzNDU2Nzg5MCJ9.ha7BcWDC2q2NI67tI9iGQ_1s6FBLbWltFafpKu02oZUhZeP55XbJ2G9pZdu2O1iCB_M7QCm4XvqTI7Q4ENa-lA_HP5VdhgfVDDmwUWpN0PPeNbU1bQ50SHC7ek-6RS-9y2LXgSeB7YzDbGG-XBxhQtZLKamWokWt3rlmmpSzujz8fhLRE4upLySHDmxkI2DUTDTLlWd8li92GhLC9exyW3mnHSW8TzwWUECNrgPdhXzu2kPQyXOW6RhYy_EoLg2_URMPQf6AAg2KHw0Q1wt1UQgcw3tNXFBRx6_-PJ0KdCIxT7DWh-mor35At0i41-y8oesRq9uzgAUZMipwKFR7xA"
	if tkn != want {
		t.Errorf("wrong token genegated got: %q", tkn)
	}
}
