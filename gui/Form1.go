// 由res2go IDE插件自动生成，不要编辑。
package gui

import (
    "github.com/ying32/govcl/vcl"
)

type TForm1 struct {
    *vcl.TForm
    Panel1       *vcl.TPanel
    Panel3       *vcl.TPanel
    GroupBox1    *vcl.TGroupBox
    Edit1        *vcl.TEdit
    GroupBox2    *vcl.TGroupBox
    Edit2        *vcl.TEdit
    StaticText1  *vcl.TStaticText
    Panel4       *vcl.TPanel
    GroupBox3    *vcl.TGroupBox
    Edit3        *vcl.TEdit
    RadioGroup1  *vcl.TRadioGroup
    RadioButton1 *vcl.TRadioButton
    RadioButton2 *vcl.TRadioButton
    Button1      *vcl.TButton
    Button2      *vcl.TButton
    StaticText2  *vcl.TStaticText
    Panel5       *vcl.TPanel
    Button3      *vcl.TButton
    Button4      *vcl.TButton
    Button5      *vcl.TButton
    Button6      *vcl.TButton
    StaticText3  *vcl.TStaticText
    ListView1    *vcl.TListView
    Panel2       *vcl.TPanel
    Panel6       *vcl.TPanel
    StaticText4  *vcl.TStaticText
    ListBox1     *vcl.TListBox
    Panel7       *vcl.TPanel
    StaticText5  *vcl.TStaticText
    ListBox2     *vcl.TListBox
    MainMenu1    *vcl.TMainMenu
    MenuItem1    *vcl.TMenuItem
    MenuItem2    *vcl.TMenuItem
    MenuItem3    *vcl.TMenuItem
    MenuItem4    *vcl.TMenuItem
    MenuItem5    *vcl.TMenuItem
    Panel8       *vcl.TPanel
    Button7      *vcl.TButton
    Button8      *vcl.TButton
    Edit4        *vcl.TEdit
    Label1       *vcl.TLabel

    //::private::
    TForm1Fields
}

var Form1 *TForm1




// vcl.Application.CreateForm(&Form1)

func NewForm1(owner vcl.IComponent) (root *TForm1)  {
    vcl.CreateResForm(owner, &root)
    return
}

var form1Bytes = []byte("\x54\x50\x46\x30\x06\x54\x46\x6F\x72\x6D\x31\x05\x46\x6F\x72\x6D\x31\x04\x4C\x65\x66\x74\x03\x5E\x01\x06\x48\x65\x69\x67\x68\x74\x03\x65\x01\x03\x54\x6F\x70\x03\xFA\x00\x05\x57\x69\x64\x74\x68\x03\x66\x02\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x27\x4A\x43\x52\x50\xE3\x80\x90\xE9\x9A\x8F\xE6\x9C\xBA\xE4\xBB\xA3\xE7\x90\x86\xE3\x80\x91\x20\x76\x33\x2E\x30\x20\x2D\x20\x62\x79\x20\x4A\x43\x30\x6F\x30\x6C\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x65\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x66\x02\x04\x4D\x65\x6E\x75\x07\x09\x4D\x61\x69\x6E\x4D\x65\x6E\x75\x31\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\xAA\x00\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\x66\x02\x05\x41\x6C\x69\x67\x6E\x07\x05\x61\x6C\x54\x6F\x70\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xAA\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x66\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x07\x54\x61\x62\x53\x74\x6F\x70\x09\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x33\x04\x4C\x65\x66\x74\x02\x01\x06\x48\x65\x69\x67\x68\x74\x03\xA8\x00\x03\x54\x6F\x70\x02\x01\x05\x57\x69\x64\x74\x68\x03\x92\x00\x05\x41\x6C\x69\x67\x6E\x07\x06\x61\x6C\x4C\x65\x66\x74\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xA8\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x92\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x09\x54\x47\x72\x6F\x75\x70\x42\x6F\x78\x09\x47\x72\x6F\x75\x70\x42\x6F\x78\x31\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x02\x39\x03\x54\x6F\x70\x02\x18\x05\x57\x69\x64\x74\x68\x03\x80\x00\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0B\xE4\xBB\xA3\xE7\x90\x86\xE6\xB1\xA0\x49\x50\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x23\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x02\x7C\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x05\x54\x45\x64\x69\x74\x05\x45\x64\x69\x74\x31\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x6F\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x04\x54\x65\x78\x74\x06\x14\x68\x74\x74\x70\x3A\x2F\x2F\x31\x30\x2E\x31\x30\x33\x2E\x39\x31\x2E\x31\x37\x39\x08\x54\x65\x78\x74\x48\x69\x6E\x74\x06\x0B\xE4\xBB\xA3\xE7\x90\x86\xE6\xB1\xA0\x49\x50\x00\x00\x00\x09\x54\x47\x72\x6F\x75\x70\x42\x6F\x78\x09\x47\x72\x6F\x75\x70\x42\x6F\x78\x32\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x02\x39\x03\x54\x6F\x70\x02\x58\x05\x57\x69\x64\x74\x68\x03\x80\x00\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0F\xE4\xBB\xA3\xE7\x90\x86\xE6\xB1\xA0\xE7\xAB\xAF\xE5\x8F\xA3\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x23\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x02\x7C\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x05\x54\x45\x64\x69\x74\x05\x45\x64\x69\x74\x32\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x03\x05\x57\x69\x64\x74\x68\x02\x50\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x04\x54\x65\x78\x74\x06\x04\x35\x30\x31\x30\x08\x54\x65\x78\x74\x48\x69\x6E\x74\x06\x0F\xE4\xBB\xA3\xE7\x90\x86\xE6\xB1\xA0\xE7\xAB\xAF\xE5\x8F\xA3\x00\x00\x00\x0B\x54\x53\x74\x61\x74\x69\x63\x54\x65\x78\x74\x0B\x53\x74\x61\x74\x69\x63\x54\x65\x78\x74\x31\x04\x4C\x65\x66\x74\x02\x01\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x02\x01\x05\x57\x69\x64\x74\x68\x03\x90\x00\x05\x41\x6C\x69\x67\x6E\x07\x05\x61\x6C\x54\x6F\x70\x09\x41\x6C\x69\x67\x6E\x6D\x65\x6E\x74\x07\x08\x74\x61\x43\x65\x6E\x74\x65\x72\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x09\x0B\x42\x6F\x72\x64\x65\x72\x53\x74\x79\x6C\x65\x07\x09\x73\x62\x73\x53\x69\x6E\x67\x6C\x65\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE4\xBB\xA3\xE7\x90\x86\xE6\xB1\xA0\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x34\x04\x4C\x65\x66\x74\x03\x93\x00\x06\x48\x65\x69\x67\x68\x74\x03\xA8\x00\x03\x54\x6F\x70\x02\x01\x05\x57\x69\x64\x74\x68\x03\x9D\x00\x05\x41\x6C\x69\x67\x6E\x07\x06\x61\x6C\x4C\x65\x66\x74\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xA8\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x9D\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x09\x54\x47\x72\x6F\x75\x70\x42\x6F\x78\x09\x47\x72\x6F\x75\x70\x42\x6F\x78\x33\x04\x4C\x65\x66\x74\x02\x05\x06\x48\x65\x69\x67\x68\x74\x02\x32\x03\x54\x6F\x70\x02\x1D\x05\x57\x69\x64\x74\x68\x02\x49\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE7\x9B\x91\xE5\x90\xAC\xE7\xAB\xAF\xE5\x8F\xA3\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x1C\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x02\x45\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x05\x54\x45\x64\x69\x74\x05\x45\x64\x69\x74\x33\x04\x4C\x65\x66\x74\x02\x06\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x03\x05\x57\x69\x64\x74\x68\x02\x30\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x04\x54\x65\x78\x74\x06\x04\x38\x30\x38\x31\x08\x54\x65\x78\x74\x48\x69\x6E\x74\x06\x0C\xE7\x9B\x91\xE5\x90\xAC\xE7\xAB\xAF\xE5\x8F\xA3\x00\x00\x00\x0B\x54\x52\x61\x64\x69\x6F\x47\x72\x6F\x75\x70\x0B\x52\x61\x64\x69\x6F\x47\x72\x6F\x75\x70\x31\x04\x4C\x65\x66\x74\x02\x05\x06\x48\x65\x69\x67\x68\x74\x02\x4D\x03\x54\x6F\x70\x02\x58\x05\x57\x69\x64\x74\x68\x02\x49\x08\x41\x75\x74\x6F\x46\x69\x6C\x6C\x09\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE4\xBB\xA3\xE7\x90\x86\xE6\xA8\xA1\xE5\xBC\x8F\x1C\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x4C\x65\x66\x74\x52\x69\x67\x68\x74\x53\x70\x61\x63\x69\x6E\x67\x02\x06\x1D\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x45\x6E\x6C\x61\x72\x67\x65\x48\x6F\x72\x69\x7A\x6F\x6E\x74\x61\x6C\x07\x18\x63\x72\x73\x48\x6F\x6D\x6F\x67\x65\x6E\x6F\x75\x73\x43\x68\x69\x6C\x64\x52\x65\x73\x69\x7A\x65\x1B\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x45\x6E\x6C\x61\x72\x67\x65\x56\x65\x72\x74\x69\x63\x61\x6C\x07\x18\x63\x72\x73\x48\x6F\x6D\x6F\x67\x65\x6E\x6F\x75\x73\x43\x68\x69\x6C\x64\x52\x65\x73\x69\x7A\x65\x1C\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x53\x68\x72\x69\x6E\x6B\x48\x6F\x72\x69\x7A\x6F\x6E\x74\x61\x6C\x07\x0E\x63\x72\x73\x53\x63\x61\x6C\x65\x43\x68\x69\x6C\x64\x73\x1A\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x53\x68\x72\x69\x6E\x6B\x56\x65\x72\x74\x69\x63\x61\x6C\x07\x0E\x63\x72\x73\x53\x63\x61\x6C\x65\x43\x68\x69\x6C\x64\x73\x12\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x4C\x61\x79\x6F\x75\x74\x07\x1D\x63\x63\x6C\x4C\x65\x66\x74\x54\x6F\x52\x69\x67\x68\x74\x54\x68\x65\x6E\x54\x6F\x70\x54\x6F\x42\x6F\x74\x74\x6F\x6D\x1B\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x43\x6F\x6E\x74\x72\x6F\x6C\x73\x50\x65\x72\x4C\x69\x6E\x65\x02\x01\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x37\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x02\x45\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x0C\x54\x52\x61\x64\x69\x6F\x42\x75\x74\x74\x6F\x6E\x0C\x52\x61\x64\x69\x6F\x42\x75\x74\x74\x6F\x6E\x31\x04\x4C\x65\x66\x74\x02\x06\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x02\x39\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE4\xBB\xA3\xE7\x90\x86\xE6\xB1\xA0\x07\x43\x68\x65\x63\x6B\x65\x64\x09\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x07\x54\x61\x62\x53\x74\x6F\x70\x09\x00\x00\x0C\x54\x52\x61\x64\x69\x6F\x42\x75\x74\x74\x6F\x6E\x0C\x52\x61\x64\x69\x6F\x42\x75\x74\x74\x6F\x6E\x32\x04\x4C\x65\x66\x74\x02\x06\x06\x48\x65\x69\x67\x68\x74\x02\x1B\x03\x54\x6F\x70\x02\x1C\x05\x57\x69\x64\x74\x68\x02\x39\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE8\x87\xAA\xE5\xAE\x9A\xE4\xB9\x89\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x31\x04\x4C\x65\x66\x74\x02\x58\x06\x48\x65\x69\x67\x68\x74\x02\x1B\x03\x54\x6F\x70\x02\x20\x05\x57\x69\x64\x74\x68\x02\x32\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x09\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x90\xAF\xE5\x8A\xA8\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x32\x04\x4C\x65\x66\x74\x02\x58\x06\x48\x65\x69\x67\x68\x74\x02\x1B\x03\x54\x6F\x70\x02\x48\x05\x57\x69\x64\x74\x68\x02\x32\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x09\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x81\x9C\xE6\xAD\xA2\x07\x45\x6E\x61\x62\x6C\x65\x64\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x00\x0B\x54\x53\x74\x61\x74\x69\x63\x54\x65\x78\x74\x0B\x53\x74\x61\x74\x69\x63\x54\x65\x78\x74\x32\x04\x4C\x65\x66\x74\x02\x01\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x02\x01\x05\x57\x69\x64\x74\x68\x03\x9B\x00\x05\x41\x6C\x69\x67\x6E\x07\x05\x61\x6C\x54\x6F\x70\x09\x41\x6C\x69\x67\x6E\x6D\x65\x6E\x74\x07\x08\x74\x61\x43\x65\x6E\x74\x65\x72\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x09\x0B\x42\x6F\x72\x64\x65\x72\x53\x74\x79\x6C\x65\x07\x09\x73\x62\x73\x53\x69\x6E\x67\x6C\x65\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE6\x8E\xA7\xE5\x88\xB6\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x04\x00\x00\x05\x54\x45\x64\x69\x74\x05\x45\x64\x69\x74\x34\x04\x4C\x65\x66\x74\x02\x58\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x82\x00\x05\x57\x69\x64\x74\x68\x02\x30\x09\x41\x6C\x69\x67\x6E\x6D\x65\x6E\x74\x07\x08\x74\x61\x43\x65\x6E\x74\x65\x72\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x05\x04\x54\x65\x78\x74\x06\x01\x35\x08\x54\x65\x78\x74\x48\x69\x6E\x74\x06\x06\x4D\x69\x6E\x4E\x75\x6D\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x50\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x69\x05\x57\x69\x64\x74\x68\x02\x46\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0F\x4D\x69\x6E\xE5\x8F\xAF\xE7\x94\xA8\xE4\xBB\xA3\xE7\x90\x86\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x35\x04\x4C\x65\x66\x74\x03\x30\x01\x06\x48\x65\x69\x67\x68\x74\x03\xA8\x00\x03\x54\x6F\x70\x02\x01\x05\x57\x69\x64\x74\x68\x03\x35\x01\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xA8\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x35\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x33\x04\x4C\x65\x66\x74\x03\xFC\x00\x06\x48\x65\x69\x67\x68\x74\x02\x1B\x03\x54\x6F\x70\x02\x1F\x05\x57\x69\x64\x74\x68\x02\x32\x07\x41\x6E\x63\x68\x6F\x72\x73\x0B\x05\x61\x6B\x54\x6F\x70\x07\x61\x6B\x52\x69\x67\x68\x74\x00\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x09\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\xAF\xBC\xE5\x85\xA5\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x34\x04\x4C\x65\x66\x74\x03\xFC\x00\x06\x48\x65\x69\x67\x68\x74\x02\x1B\x03\x54\x6F\x70\x02\x3F\x05\x57\x69\x64\x74\x68\x02\x32\x07\x41\x6E\x63\x68\x6F\x72\x73\x0B\x07\x61\x6B\x52\x69\x67\x68\x74\x00\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x09\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE4\xBF\x9D\xE5\xAD\x98\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x35\x04\x4C\x65\x66\x74\x03\xFC\x00\x06\x48\x65\x69\x67\x68\x74\x02\x1B\x03\x54\x6F\x70\x02\x5F\x05\x57\x69\x64\x74\x68\x02\x32\x07\x41\x6E\x63\x68\x6F\x72\x73\x0B\x07\x61\x6B\x52\x69\x67\x68\x74\x00\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x09\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE6\xB7\xBB\xE5\x8A\xA0\x07\x45\x6E\x61\x62\x6C\x65\x64\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x36\x04\x4C\x65\x66\x74\x03\xFC\x00\x06\x48\x65\x69\x67\x68\x74\x02\x1B\x03\x54\x6F\x70\x02\x7F\x05\x57\x69\x64\x74\x68\x02\x32\x07\x41\x6E\x63\x68\x6F\x72\x73\x0B\x07\x61\x6B\x52\x69\x67\x68\x74\x00\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x09\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x88\xA0\xE9\x99\xA4\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x00\x0B\x54\x53\x74\x61\x74\x69\x63\x54\x65\x78\x74\x0B\x53\x74\x61\x74\x69\x63\x54\x65\x78\x74\x33\x04\x4C\x65\x66\x74\x02\x01\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x02\x01\x05\x57\x69\x64\x74\x68\x03\x33\x01\x05\x41\x6C\x69\x67\x6E\x07\x05\x61\x6C\x54\x6F\x70\x09\x41\x6C\x69\x67\x6E\x6D\x65\x6E\x74\x07\x08\x74\x61\x43\x65\x6E\x74\x65\x72\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x09\x0B\x42\x6F\x72\x64\x65\x72\x53\x74\x79\x6C\x65\x07\x09\x73\x62\x73\x53\x69\x6E\x67\x6C\x65\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE8\x87\xAA\xE5\xAE\x9A\xE4\xB9\x89\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x04\x00\x00\x09\x54\x4C\x69\x73\x74\x56\x69\x65\x77\x09\x4C\x69\x73\x74\x56\x69\x65\x77\x31\x04\x4C\x65\x66\x74\x02\x01\x06\x48\x65\x69\x67\x68\x74\x03\x90\x00\x03\x54\x6F\x70\x02\x17\x05\x57\x69\x64\x74\x68\x03\xF7\x00\x05\x41\x6C\x69\x67\x6E\x07\x06\x61\x6C\x4C\x65\x66\x74\x13\x41\x75\x74\x6F\x57\x69\x64\x74\x68\x4C\x61\x73\x74\x43\x6F\x6C\x75\x6D\x6E\x09\x07\x43\x6F\x6C\x75\x6D\x6E\x73\x0E\x01\x09\x41\x6C\x69\x67\x6E\x6D\x65\x6E\x74\x07\x08\x74\x61\x43\x65\x6E\x74\x65\x72\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x02\x49\x44\x05\x57\x69\x64\x74\x68\x02\x1E\x00\x01\x09\x41\x6C\x69\x67\x6E\x6D\x65\x6E\x74\x07\x08\x74\x61\x43\x65\x6E\x74\x65\x72\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x8D\x8F\xE8\xAE\xAE\x05\x57\x69\x64\x74\x68\x02\x29\x00\x01\x09\x41\x6C\x69\x67\x6E\x6D\x65\x6E\x74\x07\x08\x74\x61\x43\x65\x6E\x74\x65\x72\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x02\x49\x50\x05\x57\x69\x64\x74\x68\x02\x6F\x00\x01\x09\x41\x6C\x69\x67\x6E\x6D\x65\x6E\x74\x07\x08\x74\x61\x43\x65\x6E\x74\x65\x72\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE7\xAB\xAF\xE5\x8F\xA3\x05\x57\x69\x64\x74\x68\x02\x41\x00\x00\x09\x47\x72\x69\x64\x4C\x69\x6E\x65\x73\x09\x09\x52\x6F\x77\x53\x65\x6C\x65\x63\x74\x09\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x05\x09\x56\x69\x65\x77\x53\x74\x79\x6C\x65\x07\x08\x76\x73\x52\x65\x70\x6F\x72\x74\x00\x00\x00\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x32\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\xBB\x00\x03\x54\x6F\x70\x03\xAA\x00\x05\x57\x69\x64\x74\x68\x03\x66\x02\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\x50\x61\x6E\x65\x6C\x32\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xBB\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x66\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x36\x04\x4C\x65\x66\x74\x02\x01\x06\x48\x65\x69\x67\x68\x74\x03\xB9\x00\x03\x54\x6F\x70\x02\x01\x05\x57\x69\x64\x74\x68\x03\xAA\x00\x05\x41\x6C\x69\x67\x6E\x07\x06\x61\x6C\x4C\x65\x66\x74\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xB9\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xAA\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x0B\x54\x53\x74\x61\x74\x69\x63\x54\x65\x78\x74\x0B\x53\x74\x61\x74\x69\x63\x54\x65\x78\x74\x34\x04\x4C\x65\x66\x74\x02\x01\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x02\x01\x05\x57\x69\x64\x74\x68\x03\xA8\x00\x05\x41\x6C\x69\x67\x6E\x07\x05\x61\x6C\x54\x6F\x70\x09\x41\x6C\x69\x67\x6E\x6D\x65\x6E\x74\x07\x08\x74\x61\x43\x65\x6E\x74\x65\x72\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x09\x0B\x42\x6F\x72\x64\x65\x72\x53\x74\x79\x6C\x65\x07\x09\x73\x62\x73\x53\x69\x6E\x67\x6C\x65\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE5\x8F\xAF\xE7\x94\xA8\xE4\xBB\xA3\xE7\x90\x86\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x08\x54\x4C\x69\x73\x74\x42\x6F\x78\x08\x4C\x69\x73\x74\x42\x6F\x78\x31\x04\x4C\x65\x66\x74\x02\x01\x06\x48\x65\x69\x67\x68\x74\x03\x81\x00\x03\x54\x6F\x70\x02\x17\x05\x57\x69\x64\x74\x68\x03\xA8\x00\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x38\x04\x4C\x65\x66\x74\x02\x01\x06\x48\x65\x69\x67\x68\x74\x02\x20\x03\x54\x6F\x70\x03\x98\x00\x05\x57\x69\x64\x74\x68\x03\xA8\x00\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x42\x6F\x74\x74\x6F\x6D\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x20\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xA8\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x37\x04\x4C\x65\x66\x74\x02\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1E\x03\x54\x6F\x70\x02\x01\x05\x57\x69\x64\x74\x68\x02\x4B\x05\x41\x6C\x69\x67\x6E\x07\x06\x61\x6C\x4C\x65\x66\x74\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE8\xBF\xBD\xE5\x8A\xA0\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x38\x04\x4C\x65\x66\x74\x02\x5C\x06\x48\x65\x69\x67\x68\x74\x02\x1E\x03\x54\x6F\x70\x02\x01\x05\x57\x69\x64\x74\x68\x02\x4B\x05\x41\x6C\x69\x67\x6E\x07\x07\x61\x6C\x52\x69\x67\x68\x74\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE8\xA6\x86\xE7\x9B\x96\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x00\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x37\x04\x4C\x65\x66\x74\x03\xAB\x00\x06\x48\x65\x69\x67\x68\x74\x03\xB9\x00\x03\x54\x6F\x70\x02\x01\x05\x57\x69\x64\x74\x68\x03\xBA\x01\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xB9\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xBA\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x0B\x54\x53\x74\x61\x74\x69\x63\x54\x65\x78\x74\x0B\x53\x74\x61\x74\x69\x63\x54\x65\x78\x74\x35\x04\x4C\x65\x66\x74\x02\x01\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x02\x01\x05\x57\x69\x64\x74\x68\x03\xB8\x01\x05\x41\x6C\x69\x67\x6E\x07\x05\x61\x6C\x54\x6F\x70\x09\x41\x6C\x69\x67\x6E\x6D\x65\x6E\x74\x07\x08\x74\x61\x43\x65\x6E\x74\x65\x72\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x09\x0B\x42\x6F\x72\x64\x65\x72\x53\x74\x79\x6C\x65\x07\x09\x73\x62\x73\x53\x69\x6E\x67\x6C\x65\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE6\x97\xA5\xE5\xBF\x97\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x08\x54\x4C\x69\x73\x74\x42\x6F\x78\x08\x4C\x69\x73\x74\x42\x6F\x78\x32\x04\x4C\x65\x66\x74\x02\x01\x06\x48\x65\x69\x67\x68\x74\x03\xA1\x00\x03\x54\x6F\x70\x02\x17\x05\x57\x69\x64\x74\x68\x03\xB8\x01\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x00\x00\x09\x54\x4D\x61\x69\x6E\x4D\x65\x6E\x75\x09\x4D\x61\x69\x6E\x4D\x65\x6E\x75\x31\x04\x4C\x65\x66\x74\x02\x5A\x03\x54\x6F\x70\x02\x10\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x09\x4D\x65\x6E\x75\x49\x74\x65\x6D\x31\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE9\x85\x8D\xE7\xBD\xAE\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x09\x4D\x65\x6E\x75\x49\x74\x65\x6D\x32\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\xAF\xBC\xE5\x85\xA5\x00\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x09\x4D\x65\x6E\x75\x49\x74\x65\x6D\x33\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\xAF\xBC\xE5\x87\xBA\x00\x00\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x09\x4D\x65\x6E\x75\x49\x74\x65\x6D\x34\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE8\xAF\xB4\xE6\x98\x8E\x00\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x09\x4D\x65\x6E\x75\x49\x74\x65\x6D\x35\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x85\xB3\xE4\xBA\x8E\x00\x00\x00\x00")

// 注册Form资源  
var _ = vcl.RegisterFormResource(Form1, &form1Bytes)
