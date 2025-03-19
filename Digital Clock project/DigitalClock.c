#include <avr/io.h>
#include <avr/interrupt.h>

volatile int seconds = 0, minutes = 0, hours = 0; //volatile because its used in the ISR and main
volatile int alarm_hours = 0, alarm_minutes = 0;
int alarm_time = 0;

// segment patterns for 7-segments display
unsigned char segment_display[] = {
    0b00111111,  // 0
    0b00000110,  // 1
    0b01011011,  // 2
    0b01001111,  // 3
    0b01100110,  // 4
    0b01101101,  // 5
    0b01111101,  // 6
    0b00000111,  // 7
    0b01111111,  // 8
    0b01101111   // 9
};

// initialize Timer1
void initialize_timer1() {
    TCCR1B = 0x0D; // compare match mode, prescaler 1024
    OCR1A = 15624; // Compare match value for 1 second --> OCR1A = clock_frequency/prescaler-1 = 16Mhz/1024-1 = 15624 (ticks to reach 1 second)
    TIMSK1 = 0x02; // OCIE1A=1 to trigger Output Compare Match A interrupt
}

ISR(TIMER1_COMPA_vect){
  if(PINA>>7 == 0){ // start the clock
	seconds++;
    if (seconds == 60){
        seconds = 0;
        minutes++;
        if (minutes == 60){
            minutes = 0;
            hours++;
            if (hours == 24){
                hours = 0;
            }
        }
   }
  }

	if (alarm_time > 0) {  // deactivate alarm after 10 seconds
		alarm_time--;
		} else {
		PORTE &= ~(1 << 7);
	}
	
}


int main(void) {
    DDRA = DDRB = DDRC = DDRD = 0x7F; // configure all pins as output except D7
    DDRE = DDRF = 0xFF; // configure all pins as output
	
    initialize_timer1();

    sei(); // enable global interrupt

while (1) {
	if (PINB >> 7 == 1) { // Set clock
		if (PINH < 24 && PINJ < 60) {
			hours = PINH;
			minutes = PINJ;
			seconds = 0;
		}
	}
	if (PIND >> 7 == 1) { // Set alarm
		if (PINH < 24 && PINJ < 60) {
			alarm_hours = PINH;
			alarm_minutes = PINJ;
		}
	}
	
	if (PINC >> 7 == 1) {  // Alarm enable
		if (alarm_hours == hours && alarm_minutes == minutes && seconds == 0) {
			PORTE |= (1 << 7);  // activate alarm
			alarm_time = 10;
		}
	}

	// Display output (time on 7-segment and alarm led)
	PORTA = segment_display[hours / 10];
	PORTB = segment_display[hours % 10];
	PORTC = segment_display[minutes / 10];
	PORTD = segment_display[minutes % 10];
	PORTE = (PORTE & (1 << 7)) | segment_display[seconds / 10];
	PORTF = segment_display[seconds % 10];
}

}