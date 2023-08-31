module github.com/jimmyaxod/scale-hackathon-webcam

go 1.20

replace signature => /home/jimmy/.config/scale/signatures/local_fetchext_v1_ce3a1b826d95bdd04852dbcf6e234bdbb33120d04e8e14e0c4d4d6018497fa97_signature/golang/host
replace github.com/loopholelabs/scale => /home/jimmy/code/scale/scale

require (
	github.com/loopholelabs/scale v0.3.20-dev11
	signature v0.1.0
)

require (
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/hashicorp/hcl/v2 v2.17.0 // indirect
	github.com/loopholelabs/polyglot v1.1.2 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/tetratelabs/wazero v1.5.0 // indirect
	github.com/zclconf/go-cty v1.13.2 // indirect
	golang.org/x/text v0.12.0 // indirect
)
