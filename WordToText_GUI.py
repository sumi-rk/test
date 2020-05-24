# coding: utf-8
import fnmatch, os, sys, win32com.client

import tkinter
from tkinter import messagebox

root = tkinter.Tk()
root.title("wordファイル→textファイル")
root.geometry("360x240")

input_label = tkinter.Label(text="変換するフォルダのパス")
input_label.place(x=10, y=70)

input_box = tkinter.Entry(width=40)
input_box.place(x=10, y=100)

def button_click():
	wa = win32com.client.gencache.EnsureDispatch("Word.Application")
	input_path = input_box.get()
	if not input_path:
		messagebox.showerror("エラー", "パスを入力してください")
	elif not os.path.isdir(input_path):
		messagebox.showerror("エラー", "存在しないパスです")
	else:
		try:
			count = 0
			# コマンドラインより探索ディレクトリpathを取得
			for path, dirs, files in os.walk(input_path): 
				for filename in files:
					# wordファイルの拡張子かをパターン・マッチング
					if not fnmatch.fnmatch(filename, "*.docx"): continue
					# wordファイルへの絶対パスを作成
					docx = os.path.abspath(os.path.join(path, filename)) 
					print("processing %s in %s" % (docx, path))
					wa.Documents.Open(docx)
					# テキストファイル保存ディレクトリ名
					txt_dirs = os.path.join(path, "old\\")
					os.makedirs(txt_dirs, exist_ok=True)
					txt = os.path.abspath(os.path.join(txt_dirs, filename[:-4] + 'txt'))
					wa.ActiveDocument.SaveAs(txt, FileFormat=win32com.client.constants.wdFormatText)
					wa.ActiveDocument.Close()
					count += 1
		except IndexError as e1:
			print("「python WordToText.py [対象フォルダ名]」で実行してください")
			print(e1)
		except Exception as e2:
			print("エラー発生")
			print(e2)
		finally:
			wa.Quit() # Wordの終了
		messagebox.showinfo("実行結果", str(count) + "個のファイルが変換されました")

button = tkinter.Button(text="実行ボタン",command=button_click)
button.place(x=10, y=130)

root.mainloop()
aa
