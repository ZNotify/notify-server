// Code generated by ent, DO NOT EDIT.

package device

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ZNotify/server/app/db/ent/generate/predicate"
	"github.com/ZNotify/server/app/manager/push/enum"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldUpdatedAt, v))
}

// Identifier applies equality check predicate on the "identifier" field. It's identical to IdentifierEQ.
func Identifier(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldIdentifier, v))
}

// Channel applies equality check predicate on the "channel" field. It's identical to ChannelEQ.
func Channel(v enum.Sender) predicate.Device {
	vc := string(v)
	return predicate.Device(sql.FieldEQ(FieldChannel, vc))
}

// ChannelMeta applies equality check predicate on the "channelMeta" field. It's identical to ChannelMetaEQ.
func ChannelMeta(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldChannelMeta, v))
}

// ChannelToken applies equality check predicate on the "channelToken" field. It's identical to ChannelTokenEQ.
func ChannelToken(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldChannelToken, v))
}

// DeviceName applies equality check predicate on the "deviceName" field. It's identical to DeviceNameEQ.
func DeviceName(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldDeviceName, v))
}

// DeviceMeta applies equality check predicate on the "deviceMeta" field. It's identical to DeviceMetaEQ.
func DeviceMeta(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldDeviceMeta, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldUpdatedAt, v))
}

// IdentifierEQ applies the EQ predicate on the "identifier" field.
func IdentifierEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldIdentifier, v))
}

// IdentifierNEQ applies the NEQ predicate on the "identifier" field.
func IdentifierNEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldIdentifier, v))
}

// IdentifierIn applies the In predicate on the "identifier" field.
func IdentifierIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldIdentifier, vs...))
}

// IdentifierNotIn applies the NotIn predicate on the "identifier" field.
func IdentifierNotIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldIdentifier, vs...))
}

// IdentifierGT applies the GT predicate on the "identifier" field.
func IdentifierGT(v string) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldIdentifier, v))
}

// IdentifierGTE applies the GTE predicate on the "identifier" field.
func IdentifierGTE(v string) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldIdentifier, v))
}

// IdentifierLT applies the LT predicate on the "identifier" field.
func IdentifierLT(v string) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldIdentifier, v))
}

// IdentifierLTE applies the LTE predicate on the "identifier" field.
func IdentifierLTE(v string) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldIdentifier, v))
}

// IdentifierContains applies the Contains predicate on the "identifier" field.
func IdentifierContains(v string) predicate.Device {
	return predicate.Device(sql.FieldContains(FieldIdentifier, v))
}

// IdentifierHasPrefix applies the HasPrefix predicate on the "identifier" field.
func IdentifierHasPrefix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasPrefix(FieldIdentifier, v))
}

// IdentifierHasSuffix applies the HasSuffix predicate on the "identifier" field.
func IdentifierHasSuffix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasSuffix(FieldIdentifier, v))
}

// IdentifierEqualFold applies the EqualFold predicate on the "identifier" field.
func IdentifierEqualFold(v string) predicate.Device {
	return predicate.Device(sql.FieldEqualFold(FieldIdentifier, v))
}

// IdentifierContainsFold applies the ContainsFold predicate on the "identifier" field.
func IdentifierContainsFold(v string) predicate.Device {
	return predicate.Device(sql.FieldContainsFold(FieldIdentifier, v))
}

// ChannelEQ applies the EQ predicate on the "channel" field.
func ChannelEQ(v enum.Sender) predicate.Device {
	vc := string(v)
	return predicate.Device(sql.FieldEQ(FieldChannel, vc))
}

// ChannelNEQ applies the NEQ predicate on the "channel" field.
func ChannelNEQ(v enum.Sender) predicate.Device {
	vc := string(v)
	return predicate.Device(sql.FieldNEQ(FieldChannel, vc))
}

// ChannelIn applies the In predicate on the "channel" field.
func ChannelIn(vs ...enum.Sender) predicate.Device {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.Device(sql.FieldIn(FieldChannel, v...))
}

// ChannelNotIn applies the NotIn predicate on the "channel" field.
func ChannelNotIn(vs ...enum.Sender) predicate.Device {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.Device(sql.FieldNotIn(FieldChannel, v...))
}

// ChannelGT applies the GT predicate on the "channel" field.
func ChannelGT(v enum.Sender) predicate.Device {
	vc := string(v)
	return predicate.Device(sql.FieldGT(FieldChannel, vc))
}

// ChannelGTE applies the GTE predicate on the "channel" field.
func ChannelGTE(v enum.Sender) predicate.Device {
	vc := string(v)
	return predicate.Device(sql.FieldGTE(FieldChannel, vc))
}

// ChannelLT applies the LT predicate on the "channel" field.
func ChannelLT(v enum.Sender) predicate.Device {
	vc := string(v)
	return predicate.Device(sql.FieldLT(FieldChannel, vc))
}

// ChannelLTE applies the LTE predicate on the "channel" field.
func ChannelLTE(v enum.Sender) predicate.Device {
	vc := string(v)
	return predicate.Device(sql.FieldLTE(FieldChannel, vc))
}

// ChannelContains applies the Contains predicate on the "channel" field.
func ChannelContains(v enum.Sender) predicate.Device {
	vc := string(v)
	return predicate.Device(sql.FieldContains(FieldChannel, vc))
}

// ChannelHasPrefix applies the HasPrefix predicate on the "channel" field.
func ChannelHasPrefix(v enum.Sender) predicate.Device {
	vc := string(v)
	return predicate.Device(sql.FieldHasPrefix(FieldChannel, vc))
}

// ChannelHasSuffix applies the HasSuffix predicate on the "channel" field.
func ChannelHasSuffix(v enum.Sender) predicate.Device {
	vc := string(v)
	return predicate.Device(sql.FieldHasSuffix(FieldChannel, vc))
}

// ChannelEqualFold applies the EqualFold predicate on the "channel" field.
func ChannelEqualFold(v enum.Sender) predicate.Device {
	vc := string(v)
	return predicate.Device(sql.FieldEqualFold(FieldChannel, vc))
}

// ChannelContainsFold applies the ContainsFold predicate on the "channel" field.
func ChannelContainsFold(v enum.Sender) predicate.Device {
	vc := string(v)
	return predicate.Device(sql.FieldContainsFold(FieldChannel, vc))
}

// ChannelMetaEQ applies the EQ predicate on the "channelMeta" field.
func ChannelMetaEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldChannelMeta, v))
}

// ChannelMetaNEQ applies the NEQ predicate on the "channelMeta" field.
func ChannelMetaNEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldChannelMeta, v))
}

// ChannelMetaIn applies the In predicate on the "channelMeta" field.
func ChannelMetaIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldChannelMeta, vs...))
}

// ChannelMetaNotIn applies the NotIn predicate on the "channelMeta" field.
func ChannelMetaNotIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldChannelMeta, vs...))
}

// ChannelMetaGT applies the GT predicate on the "channelMeta" field.
func ChannelMetaGT(v string) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldChannelMeta, v))
}

// ChannelMetaGTE applies the GTE predicate on the "channelMeta" field.
func ChannelMetaGTE(v string) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldChannelMeta, v))
}

// ChannelMetaLT applies the LT predicate on the "channelMeta" field.
func ChannelMetaLT(v string) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldChannelMeta, v))
}

// ChannelMetaLTE applies the LTE predicate on the "channelMeta" field.
func ChannelMetaLTE(v string) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldChannelMeta, v))
}

// ChannelMetaContains applies the Contains predicate on the "channelMeta" field.
func ChannelMetaContains(v string) predicate.Device {
	return predicate.Device(sql.FieldContains(FieldChannelMeta, v))
}

// ChannelMetaHasPrefix applies the HasPrefix predicate on the "channelMeta" field.
func ChannelMetaHasPrefix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasPrefix(FieldChannelMeta, v))
}

// ChannelMetaHasSuffix applies the HasSuffix predicate on the "channelMeta" field.
func ChannelMetaHasSuffix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasSuffix(FieldChannelMeta, v))
}

// ChannelMetaEqualFold applies the EqualFold predicate on the "channelMeta" field.
func ChannelMetaEqualFold(v string) predicate.Device {
	return predicate.Device(sql.FieldEqualFold(FieldChannelMeta, v))
}

// ChannelMetaContainsFold applies the ContainsFold predicate on the "channelMeta" field.
func ChannelMetaContainsFold(v string) predicate.Device {
	return predicate.Device(sql.FieldContainsFold(FieldChannelMeta, v))
}

// ChannelTokenEQ applies the EQ predicate on the "channelToken" field.
func ChannelTokenEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldChannelToken, v))
}

// ChannelTokenNEQ applies the NEQ predicate on the "channelToken" field.
func ChannelTokenNEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldChannelToken, v))
}

// ChannelTokenIn applies the In predicate on the "channelToken" field.
func ChannelTokenIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldChannelToken, vs...))
}

// ChannelTokenNotIn applies the NotIn predicate on the "channelToken" field.
func ChannelTokenNotIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldChannelToken, vs...))
}

// ChannelTokenGT applies the GT predicate on the "channelToken" field.
func ChannelTokenGT(v string) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldChannelToken, v))
}

// ChannelTokenGTE applies the GTE predicate on the "channelToken" field.
func ChannelTokenGTE(v string) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldChannelToken, v))
}

// ChannelTokenLT applies the LT predicate on the "channelToken" field.
func ChannelTokenLT(v string) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldChannelToken, v))
}

// ChannelTokenLTE applies the LTE predicate on the "channelToken" field.
func ChannelTokenLTE(v string) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldChannelToken, v))
}

// ChannelTokenContains applies the Contains predicate on the "channelToken" field.
func ChannelTokenContains(v string) predicate.Device {
	return predicate.Device(sql.FieldContains(FieldChannelToken, v))
}

// ChannelTokenHasPrefix applies the HasPrefix predicate on the "channelToken" field.
func ChannelTokenHasPrefix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasPrefix(FieldChannelToken, v))
}

// ChannelTokenHasSuffix applies the HasSuffix predicate on the "channelToken" field.
func ChannelTokenHasSuffix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasSuffix(FieldChannelToken, v))
}

// ChannelTokenEqualFold applies the EqualFold predicate on the "channelToken" field.
func ChannelTokenEqualFold(v string) predicate.Device {
	return predicate.Device(sql.FieldEqualFold(FieldChannelToken, v))
}

// ChannelTokenContainsFold applies the ContainsFold predicate on the "channelToken" field.
func ChannelTokenContainsFold(v string) predicate.Device {
	return predicate.Device(sql.FieldContainsFold(FieldChannelToken, v))
}

// DeviceNameEQ applies the EQ predicate on the "deviceName" field.
func DeviceNameEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldDeviceName, v))
}

// DeviceNameNEQ applies the NEQ predicate on the "deviceName" field.
func DeviceNameNEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldDeviceName, v))
}

// DeviceNameIn applies the In predicate on the "deviceName" field.
func DeviceNameIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldDeviceName, vs...))
}

// DeviceNameNotIn applies the NotIn predicate on the "deviceName" field.
func DeviceNameNotIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldDeviceName, vs...))
}

// DeviceNameGT applies the GT predicate on the "deviceName" field.
func DeviceNameGT(v string) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldDeviceName, v))
}

// DeviceNameGTE applies the GTE predicate on the "deviceName" field.
func DeviceNameGTE(v string) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldDeviceName, v))
}

// DeviceNameLT applies the LT predicate on the "deviceName" field.
func DeviceNameLT(v string) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldDeviceName, v))
}

// DeviceNameLTE applies the LTE predicate on the "deviceName" field.
func DeviceNameLTE(v string) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldDeviceName, v))
}

// DeviceNameContains applies the Contains predicate on the "deviceName" field.
func DeviceNameContains(v string) predicate.Device {
	return predicate.Device(sql.FieldContains(FieldDeviceName, v))
}

// DeviceNameHasPrefix applies the HasPrefix predicate on the "deviceName" field.
func DeviceNameHasPrefix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasPrefix(FieldDeviceName, v))
}

// DeviceNameHasSuffix applies the HasSuffix predicate on the "deviceName" field.
func DeviceNameHasSuffix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasSuffix(FieldDeviceName, v))
}

// DeviceNameEqualFold applies the EqualFold predicate on the "deviceName" field.
func DeviceNameEqualFold(v string) predicate.Device {
	return predicate.Device(sql.FieldEqualFold(FieldDeviceName, v))
}

// DeviceNameContainsFold applies the ContainsFold predicate on the "deviceName" field.
func DeviceNameContainsFold(v string) predicate.Device {
	return predicate.Device(sql.FieldContainsFold(FieldDeviceName, v))
}

// DeviceMetaEQ applies the EQ predicate on the "deviceMeta" field.
func DeviceMetaEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldDeviceMeta, v))
}

// DeviceMetaNEQ applies the NEQ predicate on the "deviceMeta" field.
func DeviceMetaNEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldDeviceMeta, v))
}

// DeviceMetaIn applies the In predicate on the "deviceMeta" field.
func DeviceMetaIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldDeviceMeta, vs...))
}

// DeviceMetaNotIn applies the NotIn predicate on the "deviceMeta" field.
func DeviceMetaNotIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldDeviceMeta, vs...))
}

// DeviceMetaGT applies the GT predicate on the "deviceMeta" field.
func DeviceMetaGT(v string) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldDeviceMeta, v))
}

// DeviceMetaGTE applies the GTE predicate on the "deviceMeta" field.
func DeviceMetaGTE(v string) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldDeviceMeta, v))
}

// DeviceMetaLT applies the LT predicate on the "deviceMeta" field.
func DeviceMetaLT(v string) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldDeviceMeta, v))
}

// DeviceMetaLTE applies the LTE predicate on the "deviceMeta" field.
func DeviceMetaLTE(v string) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldDeviceMeta, v))
}

// DeviceMetaContains applies the Contains predicate on the "deviceMeta" field.
func DeviceMetaContains(v string) predicate.Device {
	return predicate.Device(sql.FieldContains(FieldDeviceMeta, v))
}

// DeviceMetaHasPrefix applies the HasPrefix predicate on the "deviceMeta" field.
func DeviceMetaHasPrefix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasPrefix(FieldDeviceMeta, v))
}

// DeviceMetaHasSuffix applies the HasSuffix predicate on the "deviceMeta" field.
func DeviceMetaHasSuffix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasSuffix(FieldDeviceMeta, v))
}

// DeviceMetaEqualFold applies the EqualFold predicate on the "deviceMeta" field.
func DeviceMetaEqualFold(v string) predicate.Device {
	return predicate.Device(sql.FieldEqualFold(FieldDeviceMeta, v))
}

// DeviceMetaContainsFold applies the ContainsFold predicate on the "deviceMeta" field.
func DeviceMetaContainsFold(v string) predicate.Device {
	return predicate.Device(sql.FieldContainsFold(FieldDeviceMeta, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Device) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Device) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Device) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		p(s.Not())
	})
}
