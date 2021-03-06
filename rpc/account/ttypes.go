// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package account

import (
	"bytes"
	"fmt"
	"github.com/ShareSound/RPC-Server/rpc/shared"
	"github.com/mshockwave/thrift-go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = shared.GoUnusedProtection__
var GoUnusedProtection__ int

// Attributes:
//  - Session
//  - Email
//  - Username
type ProfileResult_ struct {
	Session  *shared.Session `thrift:"session,1,required" json:"session"`
	Email    string          `thrift:"email,2,required" json:"email"`
	Username string          `thrift:"username,3,required" json:"username"`
}

func NewProfileResult_() *ProfileResult_ {
	return &ProfileResult_{}
}

var ProfileResult__Session_DEFAULT *shared.Session

func (p *ProfileResult_) GetSession() *shared.Session {
	if !p.IsSetSession() {
		return ProfileResult__Session_DEFAULT
	}
	return p.Session
}

func (p *ProfileResult_) GetEmail() string {
	return p.Email
}

func (p *ProfileResult_) GetUsername() string {
	return p.Username
}
func (p *ProfileResult_) IsSetSession() bool {
	return p.Session != nil
}

func (p *ProfileResult_) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetSession bool = false
	var issetEmail bool = false
	var issetUsername bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetSession = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetEmail = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
			issetUsername = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetSession {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Session is not set"))
	}
	if !issetEmail {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Email is not set"))
	}
	if !issetUsername {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Username is not set"))
	}
	return nil
}

func (p *ProfileResult_) readField1(iprot thrift.TProtocol) error {
	p.Session = &shared.Session{}
	if err := p.Session.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Session), err)
	}
	return nil
}

func (p *ProfileResult_) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Email = v
	}
	return nil
}

func (p *ProfileResult_) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Username = v
	}
	return nil
}

func (p *ProfileResult_) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ProfileResult"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ProfileResult_) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("session", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:session: ", p), err)
	}
	if err := p.Session.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Session), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:session: ", p), err)
	}
	return err
}

func (p *ProfileResult_) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("email", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:email: ", p), err)
	}
	if err := oprot.WriteString(string(p.Email)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.email (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:email: ", p), err)
	}
	return err
}

func (p *ProfileResult_) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("username", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:username: ", p), err)
	}
	if err := oprot.WriteString(string(p.Username)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.username (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:username: ", p), err)
	}
	return err
}

func (p *ProfileResult_) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ProfileResult_(%+v)", *p)
}
