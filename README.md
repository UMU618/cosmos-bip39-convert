# cosmos-bip39-convert

cosmos mnemonic converter

## Convert hex string to mnemonic

```bash
$ ./cosmos-bip39-convert -hex 11e23c5687e47a7580202b006004290f811e23c5687e47a7580202b006004292
ball ball between average element depart above air about ability apart business ball ball between average element depart above air about ability apart farm
```

## Convert mnemonic to Hex formated Privatekey

```bash
$ ./cosmos-bip39-convert -mnemonic "ball ball between average element depart above air about ability apart business ball ball between average element depart above air about ability apart farm"
Bytes     : 11e23c5687e47a7580202b006004290f811e23c5687e47a7580202b00600429299
Seed      : ce27464afaf16aaa7cd4a37eee90bb93b765725f7a5908e2356708adc7e883cd173204e59b88bb7a9ea8fd262f522c3b5ec245d48d83b6f6fa7605e5f80fc5a4
Privatekey: 482eac7434983783a035062044ee76d833d9b6beebcdaa8f88f43b40a5993080
```


In Bytes: `1e23c5687e47a7580202b006004290f811e23c5687e47a7580202b006004292`**99** contains two parts,

Data    : `11e23c5687e47a7580202b006004290f811e23c5687e47a7580202b006004292`

Checksum: `99`
