#!/usr/bin/env ruby
# generate-yaml-preset.rb
require "yaml"
require "json"

$USAGE = '
  Usage: %s FILE (type|value)

  "type": generate the type definitions for a given yaml config file
  "value": generate the default value literal for a given yaml config file

  All type definitions will use the file name (including extensions) of FILE to
  generate a prefix that will precede any new type definitions created. Values
  currently set in FILE will be assumed to be intended defaults.
' % [ $0 ]

def key_to_name(key)
  key.split('-').each { |e| e.capitalize! } .join('')
end

$FILEPATH = ARGV[0]
$COMPONENT = ARGV[1]
$PREFIX = File.basename($FILEPATH)
            .split('.')
            .map { |e| key_to_name(e).capitalize() }
            .join('')

class StructDef
  attr_accessor :name, :fields

  def initialize(name, fields)
    @name = name
    @fields = []
    fields.each_pair { |k,v|
      @fields.push(Field.new(k, v))
    }
  end

  def typeLiteral()
    if @fields.empty?
      'map[string]string'
    else
      'struct {
      %s
      }' % [
        @fields.map { |f| f.fieldDefinition() } .join("\n  ")
      ]
    end
  end

  def nestedTypeDeclaration()
    out = ''
    @fields.each { |f|
      case f.default
      when StructDef
        out += "%s\n\n" % [ f.default.nestedTypeDeclaration() ]
      else
      end
    }
    out + "\n\n" + typeDeclaration()
  end

  def typeDeclaration()
    if @fields.empty?
      ''
    else
      'type %s %s' % [ @name, typeLiteral() ]
    end
  end

  def valueLiteral()
    if @fields.empty?
      'map[string]string{}'
    else
      out = "%s{\n" % [ @name ]

      @fields.each { |f|
        out += "%s: %s,\n" % [ f.name, f.valueLiteral() ]
      }

      out + "}"
    end
  end

  def as_json(options={})
    {
      name: @name,
      fields: @fields
    }
  end

  def to_json(*options)
      as_json(*options).to_json(*options)
  end
end

class Field
  attr_accessor :name, :default, :key

  def initialize(key, default)
    @name = key_to_name(key)
    case default
    when Hash
      @default = StructDef.new($PREFIX + @name, default)
    else
      @default = default
    end
    @key = key
  end

  def typeLiteral()
    case @default
    when TrueClass, FalseClass
      typeName = 'bool'
    when Integer
      typeName = 'int'
    when Float
      typeName = 'float64'
    when Array
      typeName = '[]string'
    when StructDef
      if @default.fields.empty?
        typeName = @default.typeLiteral
      else
        typeName = @default.name
      end
    else
      typeName = 'string'
    end
  end

  def fieldDefinition()
    '%s %s `yaml:"%s"`' % [@name, typeLiteral(), @key]
  end

  def valueLiteral()
    case @default
    when TrueClass, FalseClass, Integer, Float
      @default
    when Array
      '[]string{%s}' % @default.map { |e| e.to_json() } .join(',')
    when StructDef
      @default.valueLiteral()
    else
      @default.to_json()
    end
  end

  def valueField()
    "%s: %s,\n" % [ @name, valueLiteral() ]
  end

  def as_json(options={})
    {
      name: @name,
      key: @key,
      default: @default
    }
  end

  def to_json(*options)
      as_json(*options).to_json(*options)
  end
end

file = YAML.load_file($FILEPATH)
s = StructDef.new($PREFIX, file)

case $COMPONENT
when "type"
  puts s.nestedTypeDeclaration()
when "value"
  puts s.valueLiteral()
else
  puts $USAGE
end
