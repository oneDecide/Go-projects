#version 330

in vec3 vertexPosition;
in vec3 vertexNormal;
in vec2 vertexTexCoord;
in vec4 vertexColor;

uniform mat4 mvp;
uniform mat4 matModel;

out vec4 fragColor;



//our data
uniform vec4 baseColor;
uniform vec4 topColor;
uniform float time;



void main() {

    fragColor = mix(baseColor, topColor, ceil(vertexPosition.y));
    
    
    vec3 updatedPos = vertexPosition;
    
    float sway = sin(time)/2;
    updatedPos += clamp(ceil(vertexPosition.y),0,1) * vec3(sway,0.5,sway);

    gl_Position = mvp * vec4(updatedPos, 1.0);
    
}