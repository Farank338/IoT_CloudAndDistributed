#!/bin/sh
function is_in_activation {
   activation=$(systemctl status postgresql | grep "Active: active" )  
   if [ -z "$activation" ]; then
      true;
   else
      false;
   fi

   return $?;
}

while is_in_activation network;       
    do true    
    done    