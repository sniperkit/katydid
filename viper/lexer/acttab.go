package lexer

import (
	"fmt"
	"github.com/katydid/katydid/viper/token"
)

type ActionTable [NumStates]ActionRow

type ActionRow struct {
	Accept token.Type
	Ignore string
}

func (this ActionRow) String() string {
	return fmt.Sprintf("Accept=%d, Ignore=%s", this.Accept, this.Ignore)
}

var ActTab = ActionTable{
	ActionRow{ // S0
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S1
		Accept: 59,
		Ignore: "",
	},
	ActionRow{ // S2
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S3
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S4
		Accept: 34,
		Ignore: "",
	},
	ActionRow{ // S5
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S6
		Accept: 35,
		Ignore: "",
	},
	ActionRow{ // S7
		Accept: 28,
		Ignore: "",
	},
	ActionRow{ // S8
		Accept: 29,
		Ignore: "",
	},
	ActionRow{ // S9
		Accept: 41,
		Ignore: "",
	},
	ActionRow{ // S10
		Accept: 32,
		Ignore: "",
	},
	ActionRow{ // S11
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S12
		Accept: 44,
		Ignore: "",
	},
	ActionRow{ // S13
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S14
		Accept: 15,
		Ignore: "",
	},
	ActionRow{ // S15
		Accept: 15,
		Ignore: "",
	},
	ActionRow{ // S16
		Accept: 39,
		Ignore: "",
	},
	ActionRow{ // S17
		Accept: 33,
		Ignore: "",
	},
	ActionRow{ // S18
		Accept: 49,
		Ignore: "",
	},
	ActionRow{ // S19
		Accept: 27,
		Ignore: "",
	},
	ActionRow{ // S20
		Accept: 50,
		Ignore: "",
	},
	ActionRow{ // S21
		Accept: 58,
		Ignore: "",
	},
	ActionRow{ // S22
		Accept: 45,
		Ignore: "",
	},
	ActionRow{ // S23
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S24
		Accept: 37,
		Ignore: "",
	},
	ActionRow{ // S25
		Accept: 38,
		Ignore: "",
	},
	ActionRow{ // S26
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S27
		Accept: 42,
		Ignore: "",
	},
	ActionRow{ // S28
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S29
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S30
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S31
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S32
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S33
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S34
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S35
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S36
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S37
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S38
		Accept: 30,
		Ignore: "",
	},
	ActionRow{ // S39
		Accept: 36,
		Ignore: "",
	},
	ActionRow{ // S40
		Accept: 31,
		Ignore: "",
	},
	ActionRow{ // S41
		Accept: 43,
		Ignore: "",
	},
	ActionRow{ // S42
		Accept: 48,
		Ignore: "",
	},
	ActionRow{ // S43
		Accept: 8,
		Ignore: "",
	},
	ActionRow{ // S44
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S45
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S46
		Accept: 56,
		Ignore: "",
	},
	ActionRow{ // S47
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S48
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S49
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S50
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S51
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S52
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S53
		Accept: 54,
		Ignore: "",
	},
	ActionRow{ // S54
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S55
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S56
		Accept: 46,
		Ignore: "",
	},
	ActionRow{ // S57
		Accept: 17,
		Ignore: "",
	},
	ActionRow{ // S58
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S59
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S60
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S61
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S62
		Accept: 57,
		Ignore: "",
	},
	ActionRow{ // S63
		Accept: 51,
		Ignore: "",
	},
	ActionRow{ // S64
		Accept: 47,
		Ignore: "",
	},
	ActionRow{ // S65
		Accept: 52,
		Ignore: "",
	},
	ActionRow{ // S66
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S67
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S68
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S69
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S70
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S71
		Accept: 55,
		Ignore: "",
	},
	ActionRow{ // S72
		Accept: 8,
		Ignore: "",
	},
	ActionRow{ // S73
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S74
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S75
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S76
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S77
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S78
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S79
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S80
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S81
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S82
		Accept: 53,
		Ignore: "",
	},
	ActionRow{ // S83
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S84
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S85
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S86
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S87
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S88
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S89
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S90
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S91
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S92
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S93
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S94
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S95
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S96
		Accept: 59,
		Ignore: "",
	},
	ActionRow{ // S97
		Accept: 17,
		Ignore: "",
	},
	ActionRow{ // S98
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S99
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S100
		Accept: 17,
		Ignore: "",
	},
	ActionRow{ // S101
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S102
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S103
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S104
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S105
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S106
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S107
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S108
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S109
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S110
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S111
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S112
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S113
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S114
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S115
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S116
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S117
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S118
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S119
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S120
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S121
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S122
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S123
		Accept: 20,
		Ignore: "",
	},
	ActionRow{ // S124
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S125
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S126
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S127
		Accept: 17,
		Ignore: "",
	},
	ActionRow{ // S128
		Accept: 59,
		Ignore: "",
	},
	ActionRow{ // S129
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S130
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S131
		Accept: 17,
		Ignore: "",
	},
	ActionRow{ // S132
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S133
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S134
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S135
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S136
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S137
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S138
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S139
		Accept: 5,
		Ignore: "",
	},
	ActionRow{ // S140
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S141
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S142
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S143
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S144
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S145
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S146
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S147
		Accept: 25,
		Ignore: "",
	},
	ActionRow{ // S148
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S149
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S150
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S151
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S152
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S153
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S154
		Accept: 19,
		Ignore: "",
	},
	ActionRow{ // S155
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S156
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S157
		Accept: 21,
		Ignore: "",
	},
	ActionRow{ // S158
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S159
		Accept: 17,
		Ignore: "",
	},
	ActionRow{ // S160
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S161
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S162
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S163
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S164
		Accept: 10,
		Ignore: "",
	},
	ActionRow{ // S165
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S166
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S167
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S168
		Accept: 26,
		Ignore: "",
	},
	ActionRow{ // S169
		Accept: 3,
		Ignore: "",
	},
	ActionRow{ // S170
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S171
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S172
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S173
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S174
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S175
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S176
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S177
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S178
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S179
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S180
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S181
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S182
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S183
		Accept: 9,
		Ignore: "",
	},
	ActionRow{ // S184
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S185
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S186
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S187
		Accept: 11,
		Ignore: "",
	},
	ActionRow{ // S188
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S189
		Accept: 15,
		Ignore: "",
	},
	ActionRow{ // S190
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S191
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S192
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S193
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S194
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S195
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S196
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S197
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S198
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S199
		Accept: 24,
		Ignore: "",
	},
	ActionRow{ // S200
		Accept: 22,
		Ignore: "",
	},
	ActionRow{ // S201
		Accept: 23,
		Ignore: "",
	},
	ActionRow{ // S202
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S203
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S204
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S205
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S206
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S207
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S208
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S209
		Accept: 16,
		Ignore: "",
	},
	ActionRow{ // S210
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S211
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S212
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S213
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S214
		Accept: 14,
		Ignore: "",
	},
	ActionRow{ // S215
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S216
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S217
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S218
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S219
		Accept: 18,
		Ignore: "",
	},
	ActionRow{ // S220
		Accept: 12,
		Ignore: "",
	},
	ActionRow{ // S221
		Accept: 13,
		Ignore: "",
	},
	ActionRow{ // S222
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S223
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S224
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S225
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S226
		Accept: 4,
		Ignore: "",
	},
	ActionRow{ // S227
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S228
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S229
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S230
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S231
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S232
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S233
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S234
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S235
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S236
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S237
		Accept: 17,
		Ignore: "",
	},
	ActionRow{ // S238
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S239
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S240
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S241
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S242
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S243
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S244
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S245
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S246
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S247
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S248
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S249
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S250
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S251
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S252
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S253
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S254
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S255
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S256
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S257
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S258
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S259
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S260
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S261
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S262
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S263
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S264
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S265
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S266
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S267
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S268
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S269
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S270
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S271
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S272
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S273
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S274
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S275
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S276
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S277
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S278
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S279
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S280
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S281
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S282
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S283
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S284
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S285
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S286
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S287
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S288
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S289
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S290
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S291
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S292
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S293
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S294
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S295
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S296
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S297
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S298
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S299
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S300
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S301
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S302
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S303
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S304
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S305
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S306
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S307
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S308
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S309
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S310
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S311
		Accept: 0,
		Ignore: "",
	},
}