#!/usr/local/bin/perl
use AnyEvent;  
use AnyEvent::Ping;  
  
  
my $cocurrent = 100;    # ²¢·¢Êý
my @todoList = map { "192.168.6." . $_ } (1 .. 254);
  
  
my $cv = AnyEvent->condvar;  
  
my $ping = AnyEvent::Ping->new;  
  
doit() foreach 1..$cocurrent;  
sub doit{  
    my $ip = shift @todoList;  
    return if not defined $ip;  
  
    $cv->begin;  
    $ping->ping($ip, 1, sub{ done( $ip, @_ ) } );  
}  
  
sub done {  
    my ($ip, $result) = @_;  
  
    $cv->end();  
     print "Result($ip): ", $result->[0][0]," in ", $result->[0][1], " seconds\n";  
    doit();  
}  
  
$cv->recv();
