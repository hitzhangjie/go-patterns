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
	ChkTypo       bool
	ChkDirtyWords bool
}

func NewArticleProcessor(opt Options) Processor {
	p := articleProcessor{
		processors: []Processor{},
	}
	if opt.ChkTypo {
		p.processors = append(p.processors, ChkTypoProcessor{})
	}
	if opt.ChkDirtyWords {
		p.processors = append(p.processors, ChkDirtyWordsProcessor{})
	}
	return p
}

type ChkTypoProcessor struct{}

func (p ChkTypoProcessor) Process(dat []byte) ([]byte, error) {
	return dat, nil
}

type ChkDirtyWordsProcessor struct{}

func (p ChkDirtyWordsProcessor) Process(dat []byte) ([]byte, error) {
	return dat, nil
}
