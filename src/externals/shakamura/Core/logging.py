#!/usr/bin/env python3
#
# Author: Panagiotis Chartas (r0jahsm0ntar1) 
#
# This script is part of the Shakamura framework: 
# https://github.com/r0jahsm0ntar1/Shakamura

import os
from .common import system_type
from .settings import Logging_Settings

main_meta_folder = Logging_Settings.main_meta_folder_unix if system_type in ['Linux', 'Darwin'] else Logging_Settings.main_meta_folder_windows


class Shakamura_Implants_Logger:

    generated_implants_file = f'{main_meta_folder}/shakamura_generated_implants.txt'
    generated_implants_file_open = False


    @staticmethod
    def store_session_details(id, session_meta):

        try:

            while Shakamura_Implants_Logger.generated_implants_file_open:
                pass

            else:
                Shakamura_Implants_Logger.generated_implants_file_open = True
                shakamura_generated_implants = open(Shakamura_Implants_Logger.generated_implants_file, 'a')
                shakamura_generated_implants.write(f'"{id}" : {str(session_meta)}' + ',\n')
                shakamura_generated_implants.close()
                Shakamura_Implants_Logger.generated_implants_file_open = False

        except:
            pass



    @staticmethod
    def retrieve_past_sessions_data():

        if os.path.exists(Shakamura_Implants_Logger.generated_implants_file):

            try:

                while Shakamura_Implants_Logger.generated_implants_file_open:
                    pass

                else:
                    Shakamura_Implants_Logger.generated_implants_file_open = True
                    shakamura_generated_implants = open(Shakamura_Implants_Logger.generated_implants_file, 'r')
                    session_data = shakamura_generated_implants.read()
                    shakamura_generated_implants.close()
                    Shakamura_Implants_Logger.generated_implants_file_open = False
                    return '{' + session_data.strip(',\n') + '}'

            except Exception as e:
                print(e)
                pass
        
        return False
            


def clear_metadata():

    try:
        if os.path.exists(Shakamura_Implants_Logger.generated_implants_file):
            os.remove(Shakamura_Implants_Logger.generated_implants_file)
    except:
        return False
    
    return True



# Create folder to store logs and metadata
if os.path.exists(main_meta_folder):
    pass
else:
    os.makedirs(main_meta_folder)
