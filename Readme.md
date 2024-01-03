```
go-hbci: 2024/01/03 22:41:14 MessageAcknowledgement for message 1 (KA4010322411480): Code: 3060, Position: none, Text: 'Bitte beachten Sie die enthaltenen Warnungen/Hinweise.'
go-hbci: 2024/01/03 22:41:14 SegmentAcknowledgement for message 0 (), segment 4: Code: 3050, Position: none, Text: 'BPD nicht mehr aktuell, aktuelle Version enthalten.'
go-hbci: 2024/01/03 22:41:14 SegmentAcknowledgement for message 0 (), segment 4: Code: 3920, Position: none, Text: 'Zugelassene TAN-Verfahren für den Benutzer', Parameters: "962, 972, 982"
go-hbci: 2024/01/03 22:41:14 SegmentAcknowledgement for message 0 (), segment 5: Code: 3076, Position: none, Text: 'Starke Kundenauthentifizierung nicht notwendig.'
go-hbci: 2024/01/03 22:41:14 New supported security function found. Setting new security function "Smart-TAN plus manuell" (962).

SepaAccountTransactions: *fmt.wrapError:error executing HBCI request: error building request: unsupported versions [7 6 5 4]

SepaAccountBalances: *errors.errorString:error while initializing dialog: error while decrypting message: Malformed decrypted message bytes: error unmarshaling segment "HNSHK:2:4:": error unmarshaling SecurityID: *element.groupExtractor: SyntaxError at position 3: "Unexpected end of input"
("2::SItp0Gjg0IwBAAD2Qi2GhG?+owAQA")
```
