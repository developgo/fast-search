package index

import (
	"fmt"
	"testing"

	"github.com/farouqzaib/fast-search/internal/analyzer"
)

func TestInvertedIndexIndex(t *testing.T) {
	index := NewInvertedIndex()

	index.Index(1, "hello, my name is BATMAN!")
	index.Index(2, "I have come to save Gotham!")
	index.Index(3, "What is your name")

	expected := Position{DocumentID: 1, Offset: 2}

	sk := index.PostingsList["name"]
	found, err := sk.Find(expected)

	if err != nil {
		t.Fatalf("expected %v, document offset, got %v", expected, found)
	}
}

func TestInvertedIndexPrevious(t *testing.T) {
	index := NewInvertedIndex()

	index.Index(1, "hello, my name is BATMAN!")
	index.Index(2, "I have come to save Gotham!")
	index.Index(3, "What is your name")

	expected := Position{DocumentID: 1, Offset: 2}

	got, _ := index.Previous("name", Position{DocumentID: 1, Offset: 3})

	if expected.DocumentID != got.DocumentID {
		t.Fatalf("expected %v, document offset, got %v", expected, got)
	}
}

func TestInvertedIndexNext(t *testing.T) {
	index := NewInvertedIndex()

	index.Index(1, "hello, my name is BATMAN!")
	index.Index(2, "I have come to save Gotham!")
	index.Index(3, "What is your name")

	expected := Position{DocumentID: EOF, Offset: EOF}

	got, _ := index.Next("my", Position{DocumentID: 1, Offset: 1})

	fmt.Println(got)
	if expected.DocumentID != got.DocumentID && expected.Offset != got.Offset {
		t.Fatalf("expected %v, got %v", expected, got)
	}

}

func TestInvertedIndexNextPhrase(t *testing.T) {
	index := NewInvertedIndex()

	index.Index(1, "hello, my name is BATMAN!")
	index.Index(2, "I have come to save Gotham!")
	index.Index(3, "What is your name")

	expected := []Position{{DocumentID: 3, Offset: 2}, {DocumentID: 3, Offset: 3}}

	got := index.NextPhrase("your name", Position{DocumentID: BOF, Offset: BOF})

	fmt.Println(got)
	if len(got) != 2 {
		t.Fatalf("expected 2 document offsets, got %v", len(got))
	}

	if expected[1].Offset-expected[0].Offset != 1 {
		t.Fatalf("expected %v, document offset, got %v", expected, got)
	}
}

func TestInvertedIndexNextCover(t *testing.T) {
	index := NewInvertedIndex()

	index.Index(1, "hello, my name is BATMAN!")
	index.Index(2, "I have come to save Gotham!")
	index.Index(3, "What is your name")

	expected := []Position{{DocumentID: 1, Offset: 1}, {DocumentID: 1, Offset: 2}}

	tokens := analyzer.Analyze("my batman")
	got := index.NextCover(tokens, Position{DocumentID: BOF, Offset: BOF})

	b := index.Encode()
	i := index.Decode(b)

	got = i.NextCover(tokens, Position{DocumentID: BOF, Offset: BOF})

	fmt.Println(got)
	if len(got) != 2 {
		t.Fatalf("expected 2 document offsets, got %v", len(got))
	}

	if got[1].Offset-got[0].Offset != 1 {
		t.Fatalf("expected %v, document offset, got %v", expected, got)
	}
}

func TestInvertedIndexRankProximity(t *testing.T) {
	index := NewInvertedIndex()

	index.Index(1, "hello, my name is BATMAN!")
	index.Index(2, "I have come to save Gotham!")
	index.Index(3, "What is your name")
	index.Index(4, "What is my your name")

	expected := []Position{{DocumentID: 3, Offset: 2}, {DocumentID: 3, Offset: 3}}

	got := index.RankProximity("save my gotham", 10)

	if expected[1].Offset-expected[0].Offset != 1 {
		t.Fatalf("expected %v, document offset, got %v", expected, got)
	}
}

func TestInvertedIndexSerialIndex(t *testing.T) {
	index := NewInvertedIndex()

	index.Index(1, "The Project Gutenberg eBook of The Odyssey of Homer      This ebook is for the use of anyone anywhere in the United States and most other parts of the world at no cost and with almost no restrictions whatsoever. You may copy it, give it away or re-use it under the terms of the Project Gutenberg License included with this ebook or online at www.gutenberg.org. If you are not located in the United States, you will have to check the laws of the country where you are located before using this eBook.  Title: The Odyssey of Homer  Author: Homer  Translator: S. H. Butcher         Andrew Lang  Release date: April 1, 1999 [eBook #1728]                 Most recently updated: April 19, 2024  Language: English  Credits: Jim Tinsley   *** START OF THE PROJECT GUTENBERG EBOOK THE ODYSSEY OF HOMER *** The Odyssey    by Homer  DONE INTO ENGLISH PROSE  by  S. H. BUTCHER, M.A. _Fellow and Protector of University College, Oxford_ _Late Fellow of Trinity College, Cambridge_  AND  A. LANG, M.A. _Late Fellow of Merton College, Oxford_   Contents    PREFACE.  PREFACE TO THE THIRD EDITION.  INTRODUCTION.   The Odyssey  BOOK I.  BOOK II.  BOOK III.  BOOK IV.  BOOK V.  BOOK VI.  BOOK VII.  BOOK VIII.  BOOK IX.  BOOK X.  BOOK XI.  BOOK XII.  BOOK XIII.  BOOK XIV.  BOOK XV.  BOOK XVI.  BOOK XVII.  BOOK XVIII.  BOOK XIX.  BOOK XX.  BOOK XXI.  BOOK XXII.  BOOK XXIII.  BOOK XXIV.     As one that for a weary space has lain   Lulled by the song of Circe and her wine   In gardens near the pale of Proserpine, Where that Ææan isle forgets the main, And only the low lutes of love complain,   And only shadows of wan lovers pine,   As such an one were glad to know the brine Salt on his lips, and the large air again, So gladly, from the songs of modern speech   Men turn, and see the stars, and feel the free     Shrill wind beyond the close of heavy flowers,     And through the music of the languid hours They hear like Ocean on a western beach   The surge and thunder of the Odyssey.   A. L.     PREFACE.  There would have been less controversy about the proper method of Homeric translation, if critics had recognised that the question is a purely relative one, that of Homer there can be no final translation. The taste and the literary habits of each age demand different qualities in poetry, and therefore a different sort of rendering of Homer. To the men of the time of Elizabeth, Homer would have appeared bald, it seems, and lacking in ingenuity, if he had been presented in his antique simplicity. For the Elizabethan age, Chapman supplied what was then necessary, and the mannerisms that were then deemed of the essence of poetry, namely, daring and luxurious conceits. Thus in Chapman’s verse Troy must “shed her towers for tears of overthrow,” and when the winds toss Odysseus about, their sport must be called “the horrid tennis.”  In the age of Anne, “dignity” and “correctness” had to be given to Homer, and Pope gave them by aid of his dazzling rhetoric, his antitheses, his _netteté_, his command of every conventional and favourite artifice. Without Chapman’s conceits, Homer’s poems would hardly have been what the Elizabethans took for poetry; without Pope’s smoothness, and Pope’s points, the Iliad and Odyssey would have seemed rude, and harsh in the age of Anne. These great translations must always live as English poems. As transcripts of Homer they are like pictures drawn from a lost point of view. _Chaque siècle depuis le xvie a eu de ce côté son belvéder différent_. Again, when Europe woke to a sense, an almost exaggerated and certainly uncritical sense, of the value of her songs of the people, of all the ballads that Herder, Scott, Lonnrot, and the rest collected, it was commonly said that Homer was a ballad-minstrel, that the translator must imitate the simplicity, and even adopt the formulae of the ballad. Hence came the renderings of Maginn, the experiments of Mr. Gladstone, and others. There was some excuse for the error of critics who asked for a Homer in ballad rhyme. The Epic poet, the poet of gods and heroes, did indeed inherit some of the formulae of the earlier _Volks-lied_. Homer, like the author of _The Song of Roland_, like the singers of the _Kalevala_, uses constantly recurring epithets, and repeats, word for word, certain emphatic passages, messages, and so on. That custom is essential in the ballad, it is an accident not the essence of the epic. The epic is a poem of complete and elaborate art, but it still bears some birthmarks, some signs of the early popular chant, out of which it sprung, as the garden-rose springs from the wild stock, When this is recognised the demand for ballad-like simplicity and “ballad-slang” ceases to exist, and then all Homeric translations in the ballad manner cease to represent our conception of Homer. After the belief in the ballad manner follows the recognition of the romantic vein in Homer, and, as a result, came Mr. Worsley’s admirable Odyssey. This masterly translation does all that can be done for the Odyssey in the romantic style. The smoothness of the verse, the wonderful closeness to the original, reproduce all of Homer, in music and in meaning, that can be rendered in English verse. There still, however, seems an aspect Homeric poems, and a demand in connection with Homer to be recognised, and to be satisfied.  Sainte-Beuve says, with reference probably to M. Leconte de Lisle’s prose version of the epics, that some people treat the epics too much as if they were sagas. Now the Homeric epics are sagas, but then they are the sagas of the divine heroic age of Greece, and thus are told with an art which is not the art of the Northern poets. The epics are stories about the adventures of men living in most respects like the men of our own race who dwelt in Iceland, Norway, Denmark, and Sweden. The epics are, in a way, and as far as manners and institutions are concerned, historical documents. Whoever regards them in this way, must wish to read them exactly as they have reached us, without modern ornament, with nothing added or omitted. He must recognise, with Mr. Matthew Arnold, that what he now wants, namely, the simple truth about the matter of the poem, can only be given in prose, “for in a verse translation no original work is any longer recognisable.” It is for this reason that we have attempted to tell once more, in simple prose, the story of Odysseus. We have tried to transfer, not all the truth about the poem, but the historical truth, into English. In this process Homer must lose at least half his charm, his bright and equable speed, the musical current of that narrative, which, like the river of Egypt, flows from an indiscoverable source, and mirrors the temples and the palaces of unforgotten gods and kings. Without this music of verse, only a half truth about Homer can be told, but then it is that half of the truth which, at this moment, it seems most necessary to tell. This is the half of the truth that the translators who use verse cannot easily tell. They _must_ be adding to Homer, talking with Pope about “tracing the mazy lev’ret o’er the lawn,” or with Mr. Worsley about the islands that are “stars of the blue Aegaean,” or with Dr. Hawtrey about “the earth’s soft arms,” when Homer says nothing at all about the “mazy lev’ret,” or the “stars of the blue Aegaean,” or the “soft arms” of earth. It would be impertinent indeed to blame any of these translations in their place. They give that which the romantic reader of poetry, or the student of the age of Anne, looks for in verse; and without tags of this sort, a translation of Homer in verse cannot well be made to hold together.  There can be then, it appears, no final English translation of Homer. In each there must be, in addition to what is Greek and eternal, the element of what is modern, personal, and fleeting. Thus we trust that there may be room for “the pale and far-off shadow of a prose translation,” of which the aim is limited and humble. A prose translation cannot give the movement and the fire of a successful translation in verse; it only gathers, as it were, the crumbs which fall from the richer table, only tells the story, without the song. Yet to a prose translation is permitted, perhaps, that close adherence to the archaisms of the epic, which in verse become mere oddities. The double epithets, the recurring epithets of Homer, if rendered into verse, delay and puzzle the reader, as the Greek does not delay or puzzle him. In prose he may endure them, or even care to study them as the survivals of a stage of taste, which is to be found in its prime in the sagas. These double and recurring epithets of Homer are a softer form of the quaint Northern periphrases, which make the sea the “swan’s bath,” gold, the “dragon’s hoard,” men, the “ring-givers,” and so on. We do not know whether it is necessary to defend our choice of a somewhat antiquated prose. Homer has no ideas which cannot be expressed in words that are “old and plain,” and to words that are old and plain, and, as a rule, to such terms as, being used by the Translators of the Bible, are still not unfamiliar, we have tried to restrict ourselves. It may be objected, that the employment of language which does not come spontaneously to the lips, is an affectation out of place in a version of the Odyssey. To this we may answer that the Greek Epic dialect, like the English of our Bible, was a thing of slow growth and composite nature, that it was never a spoken language, nor, except for certain poetical purposes, a written language. Thus the Biblical English seems as nearly analogous to the Epic Greek, as anything that our tongue has to offer.  The few foot-notes in this book are chiefly intended to make clear some passages where there is a choice of reading. The notes at the end, which we would like to have written in the form of essays, and in company with more complete philological and archaeological studies, are chiefly meant to elucidate the life of Homer’s men.  We have received much help from many friends, and especially from Mr. R. W. Raper, Fellow of Trinity College, Oxford and Mr. Gerald Balfour, Fellow of Trinity College, Cambridge, who have aided us with many suggestions while the book was passing through the press.  In the interpretation of B. i. 411, ii. 191, v. 90, and 471, we have departed from the received view, and followed Mr. Raper, who, however, has not been able to read through the proof-sheets further than Book xii.  We have adopted La Roche’s text (Homeri Odyssea, J. La Roche, Leipzig, 1867), except in a few cases where we mention our reading in a foot-note.  The Arguments prefixed to the Books are taken, with very slight alterations, from Hobbes’ Translation of the Odyssey.  It is hoped that the Introduction added to the second edition may illustrate the growth of those national legends on which Homer worked, and may elucidate the plot of the Odyssey.    PREFACE TO THE THIRD EDITION.  We owe our thanks to the Rev. E. Warre, of Eton College, for certain corrections on nautical points. In particular, he has convinced us that the raft of Odysseus in B. v. is a raft strictly so called, and that it is not, under the poet’s description, elaborated into a ship, as has been commonly supposed. The translation of the passage (B. v. 246-261) is accordingly altered.    INTRODUCTION.")
	// index.ConcurrentIndex(2, "I have come to save Gotham!")
	// index.ConcurrentIndex(3, "What is your name")

	// expected := Position{DocumentID: 1, Offset: 2}

	// sk := index.PostingsList["name"]
	// found, err := sk.Find(expected)

	// if err != nil {
	// 	t.Fatalf("expected %v, document offset, got %v", expected, found)
	// }
}

// 0.286s
func TestInvertedIndexConcurrentIndex(t *testing.T) {
	index := NewInvertedIndex()

	tokens := analyzer.Analyze("The Project Gutenberg eBook of The Odyssey of Homer      This ebook is for the use of anyone anywhere in the United States and most other parts of the world at no cost and with almost no restrictions whatsoever. You may copy it, give it away or re-use it under the terms of the Project Gutenberg License included with this ebook or online at www.gutenberg.org. If you are not located in the United States, you will have to check the laws of the country where you are located before using this eBook.  Title: The Odyssey of Homer  Author: Homer  Translator: S. H. Butcher         Andrew Lang  Release date: April 1, 1999 [eBook #1728]                 Most recently updated: April 19, 2024  Language: English  Credits: Jim Tinsley   *** START OF THE PROJECT GUTENBERG EBOOK THE ODYSSEY OF HOMER *** The Odyssey    by Homer  DONE INTO ENGLISH PROSE  by  S. H. BUTCHER, M.A. _Fellow and Protector of University College, Oxford_ _Late Fellow of Trinity College, Cambridge_  AND  A. LANG, M.A. _Late Fellow of Merton College, Oxford_   Contents    PREFACE.  PREFACE TO THE THIRD EDITION.  INTRODUCTION.   The Odyssey  BOOK I.  BOOK II.  BOOK III.  BOOK IV.  BOOK V.  BOOK VI.  BOOK VII.  BOOK VIII.  BOOK IX.  BOOK X.  BOOK XI.  BOOK XII.  BOOK XIII.  BOOK XIV.  BOOK XV.  BOOK XVI.  BOOK XVII.  BOOK XVIII.  BOOK XIX.  BOOK XX.  BOOK XXI.  BOOK XXII.  BOOK XXIII.  BOOK XXIV.     As one that for a weary space has lain   Lulled by the song of Circe and her wine   In gardens near the pale of Proserpine, Where that Ææan isle forgets the main, And only the low lutes of love complain,   And only shadows of wan lovers pine,   As such an one were glad to know the brine Salt on his lips, and the large air again, So gladly, from the songs of modern speech   Men turn, and see the stars, and feel the free     Shrill wind beyond the close of heavy flowers,     And through the music of the languid hours They hear like Ocean on a western beach   The surge and thunder of the Odyssey.   A. L.     PREFACE.  There would have been less controversy about the proper method of Homeric translation, if critics had recognised that the question is a purely relative one, that of Homer there can be no final translation. The taste and the literary habits of each age demand different qualities in poetry, and therefore a different sort of rendering of Homer. To the men of the time of Elizabeth, Homer would have appeared bald, it seems, and lacking in ingenuity, if he had been presented in his antique simplicity. For the Elizabethan age, Chapman supplied what was then necessary, and the mannerisms that were then deemed of the essence of poetry, namely, daring and luxurious conceits. Thus in Chapman’s verse Troy must “shed her towers for tears of overthrow,” and when the winds toss Odysseus about, their sport must be called “the horrid tennis.”  In the age of Anne, “dignity” and “correctness” had to be given to Homer, and Pope gave them by aid of his dazzling rhetoric, his antitheses, his _netteté_, his command of every conventional and favourite artifice. Without Chapman’s conceits, Homer’s poems would hardly have been what the Elizabethans took for poetry; without Pope’s smoothness, and Pope’s points, the Iliad and Odyssey would have seemed rude, and harsh in the age of Anne. These great translations must always live as English poems. As transcripts of Homer they are like pictures drawn from a lost point of view. _Chaque siècle depuis le xvie a eu de ce côté son belvéder différent_. Again, when Europe woke to a sense, an almost exaggerated and certainly uncritical sense, of the value of her songs of the people, of all the ballads that Herder, Scott, Lonnrot, and the rest collected, it was commonly said that Homer was a ballad-minstrel, that the translator must imitate the simplicity, and even adopt the formulae of the ballad. Hence came the renderings of Maginn, the experiments of Mr. Gladstone, and others. There was some excuse for the error of critics who asked for a Homer in ballad rhyme. The Epic poet, the poet of gods and heroes, did indeed inherit some of the formulae of the earlier _Volks-lied_. Homer, like the author of _The Song of Roland_, like the singers of the _Kalevala_, uses constantly recurring epithets, and repeats, word for word, certain emphatic passages, messages, and so on. That custom is essential in the ballad, it is an accident not the essence of the epic. The epic is a poem of complete and elaborate art, but it still bears some birthmarks, some signs of the early popular chant, out of which it sprung, as the garden-rose springs from the wild stock, When this is recognised the demand for ballad-like simplicity and “ballad-slang” ceases to exist, and then all Homeric translations in the ballad manner cease to represent our conception of Homer. After the belief in the ballad manner follows the recognition of the romantic vein in Homer, and, as a result, came Mr. Worsley’s admirable Odyssey. This masterly translation does all that can be done for the Odyssey in the romantic style. The smoothness of the verse, the wonderful closeness to the original, reproduce all of Homer, in music and in meaning, that can be rendered in English verse. There still, however, seems an aspect Homeric poems, and a demand in connection with Homer to be recognised, and to be satisfied.  Sainte-Beuve says, with reference probably to M. Leconte de Lisle’s prose version of the epics, that some people treat the epics too much as if they were sagas. Now the Homeric epics are sagas, but then they are the sagas of the divine heroic age of Greece, and thus are told with an art which is not the art of the Northern poets. The epics are stories about the adventures of men living in most respects like the men of our own race who dwelt in Iceland, Norway, Denmark, and Sweden. The epics are, in a way, and as far as manners and institutions are concerned, historical documents. Whoever regards them in this way, must wish to read them exactly as they have reached us, without modern ornament, with nothing added or omitted. He must recognise, with Mr. Matthew Arnold, that what he now wants, namely, the simple truth about the matter of the poem, can only be given in prose, “for in a verse translation no original work is any longer recognisable.” It is for this reason that we have attempted to tell once more, in simple prose, the story of Odysseus. We have tried to transfer, not all the truth about the poem, but the historical truth, into English. In this process Homer must lose at least half his charm, his bright and equable speed, the musical current of that narrative, which, like the river of Egypt, flows from an indiscoverable source, and mirrors the temples and the palaces of unforgotten gods and kings. Without this music of verse, only a half truth about Homer can be told, but then it is that half of the truth which, at this moment, it seems most necessary to tell. This is the half of the truth that the translators who use verse cannot easily tell. They _must_ be adding to Homer, talking with Pope about “tracing the mazy lev’ret o’er the lawn,” or with Mr. Worsley about the islands that are “stars of the blue Aegaean,” or with Dr. Hawtrey about “the earth’s soft arms,” when Homer says nothing at all about the “mazy lev’ret,” or the “stars of the blue Aegaean,” or the “soft arms” of earth. It would be impertinent indeed to blame any of these translations in their place. They give that which the romantic reader of poetry, or the student of the age of Anne, looks for in verse; and without tags of this sort, a translation of Homer in verse cannot well be made to hold together.  There can be then, it appears, no final English translation of Homer. In each there must be, in addition to what is Greek and eternal, the element of what is modern, personal, and fleeting. Thus we trust that there may be room for “the pale and far-off shadow of a prose translation,” of which the aim is limited and humble. A prose translation cannot give the movement and the fire of a successful translation in verse; it only gathers, as it were, the crumbs which fall from the richer table, only tells the story, without the song. Yet to a prose translation is permitted, perhaps, that close adherence to the archaisms of the epic, which in verse become mere oddities. The double epithets, the recurring epithets of Homer, if rendered into verse, delay and puzzle the reader, as the Greek does not delay or puzzle him. In prose he may endure them, or even care to study them as the survivals of a stage of taste, which is to be found in its prime in the sagas. These double and recurring epithets of Homer are a softer form of the quaint Northern periphrases, which make the sea the “swan’s bath,” gold, the “dragon’s hoard,” men, the “ring-givers,” and so on. We do not know whether it is necessary to defend our choice of a somewhat antiquated prose. Homer has no ideas which cannot be expressed in words that are “old and plain,” and to words that are old and plain, and, as a rule, to such terms as, being used by the Translators of the Bible, are still not unfamiliar, we have tried to restrict ourselves. It may be objected, that the employment of language which does not come spontaneously to the lips, is an affectation out of place in a version of the Odyssey. To this we may answer that the Greek Epic dialect, like the English of our Bible, was a thing of slow growth and composite nature, that it was never a spoken language, nor, except for certain poetical purposes, a written language. Thus the Biblical English seems as nearly analogous to the Epic Greek, as anything that our tongue has to offer.  The few foot-notes in this book are chiefly intended to make clear some passages where there is a choice of reading. The notes at the end, which we would like to have written in the form of essays, and in company with more complete philological and archaeological studies, are chiefly meant to elucidate the life of Homer’s men.  We have received much help from many friends, and especially from Mr. R. W. Raper, Fellow of Trinity College, Oxford and Mr. Gerald Balfour, Fellow of Trinity College, Cambridge, who have aided us with many suggestions while the book was passing through the press.  In the interpretation of B. i. 411, ii. 191, v. 90, and 471, we have departed from the received view, and followed Mr. Raper, who, however, has not been able to read through the proof-sheets further than Book xii.  We have adopted La Roche’s text (Homeri Odyssea, J. La Roche, Leipzig, 1867), except in a few cases where we mention our reading in a foot-note.  The Arguments prefixed to the Books are taken, with very slight alterations, from Hobbes’ Translation of the Odyssey.  It is hoped that the Introduction added to the second edition may illustrate the growth of those national legends on which Homer worked, and may elucidate the plot of the Odyssey.    PREFACE TO THE THIRD EDITION.  We owe our thanks to the Rev. E. Warre, of Eton College, for certain corrections on nautical points. In particular, he has convinced us that the raft of Odysseus in B. v. is a raft strictly so called, and that it is not, under the poet’s description, elaborated into a ship, as has been commonly supposed. The translation of the passage (B. v. 246-261) is accordingly altered.    INTRODUCTION.")
	index.ConcurrentIndex(1, tokens)
	// index.ConcurrentIndex(2, "I have come to save Gotham!")
	// index.ConcurrentIndex(3, "What is your name")

	// expected := Position{DocumentID: 1, Offset: 2}

	// sk := index.PostingsList["name"]
	// found, err := sk.Find(expected)

	// if err != nil {
	// 	t.Fatalf("expected %v, document offset, got %v", expected, found)
	// }
}
