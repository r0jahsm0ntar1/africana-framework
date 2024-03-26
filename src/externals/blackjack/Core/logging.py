#!/usr/bin/env python3
#
# Author: Rojahs Montari (r0jahsm0ntar1) 
#
# This script is part of the Blackjack framework: 
# https://github.com/r0jahsm0ntar1/africana-framework

import os
from .common import system_type
from .settings import Logging_Settings

main_meta_folder = Logging_Settings.main_meta_folder_unix if system_type in ['Linux', 'Darwin'] else Logging_Settings.main_meta_folder_windows


class BlackJack_Implants_Logger:

    generated_implants_file = f'{main_meta_folder}/blackjack_generated_implants.txt'
    generated_implants_file_open = False


    @staticmethod
    def store_session_details(id, session_meta):

        try:

            while BlackJack_Implants_Logger.generated_implants_file_open:
                pass

            else:
                BlackJack_Implants_Logger.generated_implants_file_open = True
                blackjack_generated_implants = open(BlackJack_Implants_Logger.generated_implants_file, 'a')
                blackjack_generated_implants.write(f'"{id}" : {str(session_meta)}' + ',\n')
                blackjack_generated_implants.close()
                BlackJack_Implants_Logger.generated_implants_file_open = False

        except:
            pass



    @staticmethod
    def retrieve_past_sessions_data():

        if os.path.exists(BlackJack_Implants_Logger.generated_implants_file):

            try:

                while BlackJack_Implants_Logger.generated_implants_file_open:
                    pass

                else:
                    BlackJack_Implants_Logger.generated_implants_file_open = True
                    blackjack_generated_implants = open(BlackJack_Implants_Logger.generated_implants_file, 'r')
                    session_data = blackjack_generated_implants.read()
                    blackjack_generated_implants.close()
                    BlackJack_Implants_Logger.generated_implants_file_open = False
                    return '{' + session_data.strip(',\n') + '}'

            except Exception as e:
                print(e)
                pass
        
        return False
            


def clear_metadata():

    try:
        if os.path.exists(BlackJack_Implants_Logger.generated_implants_file):
            os.remove(BlackJack_Implants_Logger.generated_implants_file)
    except:
        return False
    
    return True



# Create folder to store logs and metadata
if os.path.exists(main_meta_folder):
    pass
else:
    os.makedirs(main_meta_folder)
