<h1>Body temperature</h1>
A change in body temperature is one of the first signs of developing disease such as cold or influenza. Even in itself, too high or too low a body temperature may be really dangerous.<br>
<br>
We recognise four state of body temperature:<br>
<ul>
	<li>hypothermia (below 35.0ºC),</li>
	<li>normal (at least 35.0ºC and at most 37.5ºC),</li>
	<li>fever (above 37.5ºC and at most 40.0ºC),</li>
	<li>hyperpyrexia (above 40.0ºC).</li>
</ul>
Given a measurement, return the patient's body temperature state.<br>
<br>
write a function:<br>
<br>
<code>func temperature(T string) string</code>

that, given a string T, describing a patient's body temperature in the Celsius scale, returns the name of the patient's temperature state. The string T is the form "DD.D", where D denotes digit.<br>
<br>
<b>Examples:</b><br>
<br>
1. Given T = "34.5", the function should return "hypothermia".<br>
2. Given T = "35.0", the function should return "normal".<br>
3. Given T = "37.6", the function should return "fever".<br>
4. Given T = "41.0", the function should return "hyperpyrexia".<br>
<br>
Assume that:<br>
<ul>
	<li>string T is in format "DD.D", where D denotes a digit;</li>
	<li>string T describes temperature within the range [34.0ºC, 42.0ºC]</li>
</ul>


