#!/usr/bin/python3

import argparse
import sys

class GWF:
	def __init__(self, action=None, **kwargs):
		"""
		Provides a tool to automatically generate
		elements of GWF.
		"""
		self.routeFile = "./routes.go"
		self.controllerDir = "./src/controllers/"
		self.viewDir = "./src/views/"
		self.modelDir = "./src/models/"

		if action == "generate":
			self.generate(kwargs=kwargs),
		elif action == "map":
			self.map(),
		else:
			self.help()
	
	def help(self):
		print("help page")

	def generate(self, **kwargs):
		print("generating")

	def map(self):
		print("mapping")

if __name__ == "__main__":
	# Get CLI arguments.
	parser = argparse.ArgumentParser()

	parser.add_argument("--action", choices=["map", "generate"] ,help="An action to be executed by GWFtool.")
	subparsers = parser.add_subparsers()

	# create the parser for the "generate" command
	parser_generate = subparsers.add_parser("generate", help="Sub command help for generate")
	parser_generate.add_argument("url", help="The path that it should be accessed on.")

	parser.parse_args(sys.argv)
	gwf = GWF()