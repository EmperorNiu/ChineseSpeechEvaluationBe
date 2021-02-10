# from docx import Document
# from docx import table

# doc_path = './report_template3.docx'
# doc = Document(doc_path)
# print(doc.tables[0].rows[0].cells[0].text)
# print(doc.tables[0].rows[0].cells[6].text)
# print(doc.tables[0].rows[0].cells[11].text)
# print(doc.tables[0].rows[0].cells[12].text)

# print(doc.tables[0].rows[6].cells[0].text)
# print(doc.tables[0].rows[6].cells[1].text)
# print(doc.tables[0].rows[6].cells[2].text)
# print(doc.tables[0].rows[6].cells[3].text)
# print(doc.tables[0].rows[6].cells[5].text)
# print(doc.tables[0].rows[6].cells[7].text)
# print(doc.tables[0].rows[6].cells[8].text)
# print(doc.tables[0].rows[6].cells[9].text)
# print(doc.tables[0].rows[6].cells[10].text)
# print(doc.tables[0].rows[6].cells[11].text)
# print(doc.tables[0].rows[6].cells[12].text)
# print(doc.tables[0].rows[6].cells[13].text)
# print(doc.tables[0].rows[22].cells[4].text)

from pypinyin import pinyin

tmp = pinyin('朝阳')
tmp2 = []
for i in tmp:
    tmp2.append(i[0])
tmp = ' '.join(tmp2)
print(tmp)
