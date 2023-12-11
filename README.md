# otbias(1) - calculate anode current and dissipation using the output transformer method

Kozmix Go, 18-JUN-2023

```
otbias Vcentre Vanode Rtransformer [Panode]
```


<a name="description"></a>

# Description

You use
**otbias**
to calculate anode current and dissipation of a power valve by
measuring voltages on the output transformer. It takes three or four
arguments:


<a name="options"></a>

# Options


* **Vcentre**  
  Voltage from the output transformer centre tap to ground, in Volts.
* **Vanode**  
  Voltage from the anode to ground (this can usually be measured from
  the transformer as well as from the correct pin on the valve socket),
  in Volts.
  

Both voltages are measured on a
_live_
amplifier that has had a couple of minutes to warm up and stabilise,
and with the volume at 0.


* **Rtransformer**  
  DC resistance in Ohms between the centre tap and the anode, meansured
  with the amplifier turned
  _off_
  (obviously). You might want to drain the filter capacitors before
  doing this.
  

Optionally followed by:


* **Panode**  
  desired anode power dissipation in Watt. This will calculate the
  voltage drop between centre tap and anode corresponding to this anode
  current, for easy biasing.
  
  Note that this is somewhat meaningless since changing the bias to
  achieve this voltage drop is quite likely to change the anode voltage
  which in its turn changes this calculation.
  
  

<a name="example"></a>

# Example

.EX
$ ./otbias 250 247.76 109 10
centre tap voltage            = 250 V
anode voltage                 = 247.8 V
output transformer resistance = 109 Ω
desired anode dissipation     = 10 W

voltage drop                  = 2.24 V
anode current                 = 20.55 mA
anode dissipation             = 5.092 W

desired voltage drop          = 4.399 V
.EE


<a name="author"></a>

# Author

svm

