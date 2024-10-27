package decorator

// Supposing you want to submit an article to media, after submiting
// their system may process your article step by step.
//
// For example, check dirty words, check typo, check hate words, ...
// Here, we define Processor to represent any processing step.

type Processor interface {
	Process([]byte) ([]byte, error)
}

type articleProcessor struct {
	processors []Processor
	opts       Options
}

func (p articleProcessor) Process(dat []byte) ([]byte, error) {
	for _, pp := range p.processors {
		v, err := pp.Process(dat)
		if err != nil {
			return nil, err
		}
		dat = v
	}
	return dat, nil
}

type Options struct {
	checkTypo       bool
	checkDirtyWords bool
}

func NewArticleProcessor(opt Options) Processor {
	p := articleProcessor{
		processors: []Processor{},
	}
	if opt.checkTypo {
		p.processors = append(p.processors, checkTypoProcessor{})
	}
	if opt.checkDirtyWords {
		p.processors = append(p.processors, checkDirtyWordsProcessor{})
	}
	return p
}

type checkTypoProcessor struct{}

func (p checkTypoProcessor) Process(dat []byte) ([]byte, error) {
	return dat, nil
}

type checkDirtyWordsProcessor struct{}

func (p checkDirtyWordsProcessor) Process(dat []byte) ([]byte, error) {
	return dat, nil
}

// well, here there's no check hate words processor, which will be defined
// when we want to decorate our articleProcessor in nohateArticleProcessor.
